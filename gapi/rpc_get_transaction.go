package gapi

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateGetTransaction(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	parsedUUID, _ := uuid.Parse(req.GetId())

	arg := db.GetTransactionParams{
		ID:        parsedUUID,
		AccountID: authPayload.AccountId,
	}
	transaction, err := server.store.GetTransaction(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "transaction not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get transaction: %s", err)
	}

	rsp := &pb.GetTransactionResponse{
		Transaction: convertTransaction(transaction),
	}
	return rsp, nil
}

func validateGetTransaction(req *pb.GetTransactionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUUID(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
