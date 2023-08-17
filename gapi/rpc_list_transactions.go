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

func (server *Server) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateListTransactionsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListTransactionsByAccountParams{
		AccountID: authPayload.AccountId,
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	}

	transactions, err := server.store.ListTransactionsByAccount(ctx, arg)
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

func validateListTransactionsRequest(req *pb.ListTransactionsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateNumberLimit(req.GetLimit(), 0, 200); err != nil {
		violations = append(violations, fieldViolation("limit", err))
	}
	if err := val.ValidateNumberLimit(req.GetOffset(), 0, 200); err != nil {
		violations = append(violations, fieldViolation("offset", err))
	}

	return violations
}
