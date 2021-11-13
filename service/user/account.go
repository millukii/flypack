package user

import (
	"context"
	"flypack/models"
	"flypack/repository"
	"fmt"
)

type UserAccountService interface {
	RegisterNewAccount(ctx context.Context, req *models.RegisterNewUserRequest, passwordGenerator PasswordGenerator) (*models.RegisterNewUserResponse, error)
	ResetPassword(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	ActivateAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	DeActivateAccount(ctx context.Context, req *models.RegisterNewUserRequest) (*models.RegisterNewUserResponse, error)
	Login(ctx context.Context, req *models.RequestLogin)(*models.LoginResponse, error)
}

type userAccount struct{
	userRepository repository.UserRepository
}

func NewUserAccountService(repository repository.UserRepository) (UserAccountService){

	newService := &userAccount{
		userRepository: repository,
	}

	return newService
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

func (user userAccount) RegisterNewAccount(ctx context.Context, req *models.RegisterNewUserRequest, passwordGenerator PasswordGenerator) (*models.RegisterNewUserResponse, error){

	secureAndOneTimePassword, err := passwordGenerator.GeneratePassword()
	if err != nil {
		fmt.Println("Error secureAndOneTimePassword the value now is empty")
	}
	newUser := &models.User{
		ID: 0,
		User: req.User,
		Register: req.Register,
		Role: req.Role,
		Password: secureAndOneTimePassword ,
		State: 1,
	}

	created, err := user.userRepository.CreateUser(ctx, newUser)

	if err != nil {
		return nil, err
	}
	
	return &models.RegisterNewUserResponse{
		User: created.User,
		Role: created.Role,
		Message: "Usuario creado satisfactoriamente",
	}, nil
}

func (user userAccount)	Login(ctx context.Context, req *models.RequestLogin)(*models.LoginResponse, error){
	return nil,nil
}