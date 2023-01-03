package storage

import (
	"MyProjects/RentCar_gRPC/auth_rentcar_service/protogen/blogpost"
)

type StorageInter interface {
	AddNewUser(id string, box *blogpost.CreateUserRequest) error
	GetUserById(id string) (*blogpost.User, error)
	GetUserList(offset, limit int, search string) (dataset *blogpost.GetUserListResponse, err error)
	UpdateUser(box *blogpost.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUserByUsername(username string) (*blogpost.User, error)
}
