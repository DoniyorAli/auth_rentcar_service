package authorization

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/authorization"
	"MyProjects/RentCar_gRPC/auth_rentcar_service/security"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authService) Login(ctx context.Context, req *authorization.LoginRequest) (*authorization.TokenResponse, error) {
	log.Println("Login...")

	errAuth := errors.New("username or password wrong")

	user, err := s.stg.GetUserByUsername(req.Username)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Errorf(codes.Unauthenticated, errAuth.Error())
	}

	match, err := security.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "security.ComparePassword: %s", err.Error())
	}

	if !match {
		return nil, status.Errorf(codes.Unauthenticated, errAuth.Error())
	}

	m := map[string]interface{}{
		"user_id":  user.Id,
		"username": user.Username,
	}

	tokenStr, err := security.GenerateJWT(m, 10*time.Minute, s.cfg.SecretKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "security.GenerateJWT: %s", err.Error())
	}

	return &authorization.TokenResponse{
		Token: tokenStr,
	}, nil
}

// HasAccess ...
func (s *authService) HasAccess(ctx context.Context, req *authorization.TokenRequest) (*authorization.HasAccessResponse, error) {
	log.Println("HasAccess...")

	result, err := security.ParseClaims(req.Token, s.cfg.SecretKey)
	if err != nil {
		log.Println(status.Errorf(codes.Unauthenticated, "security.ParseClaims: %s", err.Error()))
		return &authorization.HasAccessResponse{
			User:      nil,
			HasAccess: false,
		}, nil
	}

	log.Println(result.Username)

	user, err := s.stg.GetUserById(result.UserID)
	if err != nil {
		log.Println(status.Errorf(codes.Unauthenticated, "s.stg.GetUserByID: %s", err.Error()))
		return &authorization.HasAccessResponse{
			User:      nil,
			HasAccess: false,
		}, nil
	}

	return &authorization.HasAccessResponse{
		User:      user,
		HasAccess: true,
	}, nil
}