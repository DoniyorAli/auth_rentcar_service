package storage

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/authorization"
)

type StorageInter interface {
	AddNewUser(id string, box *authorization.CreateUserRequest) error
	GetUserById(id string) (*authorization.User, error)
	GetUserList(offset, limit int, search string) (dataset *authorization.GetUserListResponse, err error)
	UpdateUser(box *authorization.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUserByUsername(username string) (*authorization.User, error)
}
