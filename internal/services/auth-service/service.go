package auth_service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	"vote-ui/internal/pb/auth_service"
)

type AuthService struct {
	authService auth_service.AuthServiceClient
}

func NewAuthService(grpcClient *grpc.ClientConn) *AuthService {
	return &AuthService{
		authService: auth_service.NewAuthServiceClient(grpcClient),
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*User, error) {
	resp, err := s.authService.Auth(ctx, &auth_service.AuthRequest{
		Login: username,
		Pass:  password,
	})
	if err != nil {
		return nil, fmt.Errorf("rpc.Auth: %w", err)
	}
	return &User{
		Token: &Token{
			Token:     resp.GetUser().GetToken().GetToken(),
			ExpiresAt: time.Unix(resp.GetUser().GetToken().GetExpires(), 0),
		},
		ID: resp.GetUser().GetUserID(),
	}, nil
}
