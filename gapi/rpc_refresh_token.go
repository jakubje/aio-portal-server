package gapi

import (
	"context"
	"errors"
	"fmt"
	db "github.com/jakub/aioportal/server/db/sqlc"
	"github.com/jakub/aioportal/server/pb"
	"github.com/jakub/aioportal/server/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (server *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	violations := validateRefreshTokenRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to has verify token: %s", err)
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "failed to find token: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get token: %s", err)
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	if session.Email != refreshPayload.Email {
		err := fmt.Errorf("incorrect session user")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	if session.RefreshToken != req.GetRefreshToken() {
		err := fmt.Errorf("mismatch session token")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return nil, status.Errorf(codes.Unauthenticated, "%s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.AccountId,
		session.Email,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err)
	}

	rsp := &pb.RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return rsp, nil

}

func validateRefreshTokenRequest(req *pb.RefreshTokenRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetRefreshToken(), 4, 5000); err != nil {
		violations = append(violations, fieldViolation("refresh_token", err))
	}

	return violations
}
