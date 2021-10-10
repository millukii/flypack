package user

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type UserAccountService interface {
	RegisterNewAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	ResetPassword(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	ActivateAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	DeActivateAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)

}

type userAccount struct{
	userRepository *repository.UserRepository
}

func NewUserAccountService(repository *repository.UserRepository) (UserAccountService, error){

	newService := &userAccount{
		userRepository: repository,
	}

	return newService, nil
}

func (user userAccount) 	ResetPassword(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error){
	return nil, nil
}
func (user userAccount) ActivateAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error){
	return nil, nil
}
func (user userAccount) DeActivateAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error){
	return nil, nil
}

func (user userAccount) RegisterNewAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error){
	return nil, nil
}
