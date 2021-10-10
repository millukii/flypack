package user

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type UserService interface {
	EditUserInfo(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	GetUserInfo(ctx context.Context, id string) (*models.UserListView, error)
	GetAllUsersInfo(ctx context.Context, req *models.GetAllUserRequest) (*models.	AllUserResponse, error)

}

type user struct{
	userRepository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) (UserService, error){

	newService := &user{
		userRepository: repository,
	}

	return newService, nil
}

func (user user) EditUserInfo(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error){
	return nil, nil
}
func (user user) GetUserInfo(ctx context.Context, id string) (*models.UserListView, error){
	return nil, nil
}
func (user user) 	GetAllUsersInfo(ctx context.Context, req *models.GetAllUserRequest) (*models.	AllUserResponse, error){
	return nil, nil
}