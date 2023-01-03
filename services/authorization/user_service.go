package authorization

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/config"
	blogpost "MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/blogpost"
	"MyProjects/RentCar_gRPC/auth_rentcar_service/security"
	"MyProjects/RentCar_gRPC/auth_rentcar_service/storage"
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	cfg config.Config
	stg storage.StorageInter
	blogpost.UnimplementedAuthServiceServer
}

func NewAuthService(cfg config.Config, stg storage.StorageInter) *authService {
	return &authService{
		cfg: cfg,
		stg: stg,
	}
}

func (s *authService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong{
		Message: "OK",
	}, nil
}

//?==============================================================================================================

func (s *authService) CreateUser(ctx context.Context, req *blogpost.CreateUserRequest) (*blogpost.User, error) {
	id := uuid.New()

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "security.HashPassword: %s", err.Error())
	}

	req.Password = hashedPassword

	err = s.stg.AddNewUser(id.String(), req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddNewUser: %s", err.Error())
	}

	user, err := s.stg.GetUserById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	return user, nil
}

//?==============================================================================================================

func (s *authService) UpdateUser(ctx context.Context, req *blogpost.UpdateUserRequest) (*blogpost.User, error) {

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "security.HashPassword: %s", err.Error())
	}

	req.Password = hashedPassword

	err = s.stg.UpdateUser(req)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateUser: %s", err.Error())
	}

	user, err := s.stg.GetUserById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	return user, nil
}

//?==============================================================================================================

func (s *authService) DeleteUser(ctx context.Context, req *blogpost.DeleteUserRequest) (*blogpost.User, error) {

	user, err := s.stg.GetUserById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	err = s.stg.DeleteUser(user.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteUser: %s", err.Error())
	}

	return user, nil
}

//?==============================================================================================================

func (s *authService) GetUserList(ctx context.Context, req *blogpost.GetUserListRequest) (*blogpost.GetUserListResponse, error) {
	res, err := s.stg.GetUserList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserList: %s", err.Error())
	}

	return res, nil
}

//?==============================================================================================================

func (s *authService) GetUserByID(ctx context.Context, req *blogpost.GetUserByIDRequest) (*blogpost.User, error) {
	user, err := s.stg.GetUserById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetUserById: %s", err.Error())
	}

	return user, nil
}
