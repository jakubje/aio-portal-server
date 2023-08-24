package gapi

import (
	"context"
	"github.com/jackc/pgx/v5"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListTransactionsByCoin(ctx context.Context, req *pb.ListTransactionsByCoinRequest) (*pb.ListTransactionsResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateListTransactionsByCoinRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListTransactionsByAccountByCoinParams{
		AccountID: authPayload.AccountId,
		Symbol:    req.GetSymbol(),
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	}

	transactions, err := server.store.ListTransactionsByAccountByCoin(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "no transactions found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get transactions: %s", err)
	}

	rsp := &pb.ListTransactionsResponse{
		Total:        int64(len(transactions)),
		Transactions: convertTransactions(transactions),
	}
	return rsp, nil
}

func validateListTransactionsByCoinRequest(req *pb.ListTransactionsByCoinRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetSymbol(), 1, 6); err != nil {
		violations = append(violations, fieldViolation("symbol", err))
	}
	if err := val.ValidateNumberLimit(req.GetLimit(), 0, 200); err != nil {
		violations = append(violations, fieldViolation("limit", err))
	}
	if err := val.ValidateNumberLimit(req.GetOffset(), 0, 200); err != nil {
		violations = append(violations, fieldViolation("offset", err))
	}

	return violations
}
