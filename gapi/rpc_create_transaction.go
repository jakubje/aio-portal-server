package gapi

import (
	"context"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

func (server *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateTransactionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	coinToAddDetails, err := server.store.GetCoin(ctx, req.GetSymbol())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add this coin: %s", err)
	}
	pricePerCoin, err := strconv.ParseFloat(coinToAddDetails.Price, 8)
	quantity := req.GetAmount() / pricePerCoin

	arg := db.CreateTransactionParams{
		AccountID:      authPayload.AccountId,
		PortfolioID:    req.GetPortfolioId(),
		Symbol:         req.GetSymbol(),
		Type:           req.GetType(),
		Amount:         req.GetAmount(),
		PricePerCoin:   pricePerCoin,
		Quantity:       quantity,
		TimeTransacted: time.Now(),
		TimeCreated:    time.Now(),
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transaction: %s", err)
	}

	rsp := &pb.CreateTransactionResponse{
		Transaction: convertTransaction(transaction),
	}
	return rsp, nil
}

func validateCreateTransactionRequest(req *pb.CreateTransactionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetPortfolioId()); err != nil {
		violations = append(violations, fieldViolation("portfolio_id", err))
	}
	if err := val.ValidateString(req.GetSymbol(), 1, 6); err != nil {
		violations = append(violations, fieldViolation("symbol", err))
	}
	if err := val.ValidateType(req.GetType()); err != nil {
		violations = append(violations, fieldViolation("type", err))
	}
	//if err := val.ValidateFloat(req.GetPricePerCoin()); err != nil {
	//	violations = append(violations, fieldViolation("price_per_coin", err))
	//}
	//if err := val.ValidateFloat(req.GetQuantity()); err != nil {
	//	violations = append(violations, fieldViolation("quantity", err))
	//}
	return violations
}
