package user

import (
	"context"
	"flypack/models"
	"flypack/repository"
	"fmt"

)

type UserService interface {
	EditUserInfo(ctx context.Context,id string, req *models.RegisterUpdateUserRequest) (*models.RegisterNewUserResponse, error)
	GetUserInfo(ctx context.Context, id string) (*models.UserListView, error)
	GetAllUsersInfo(ctx context.Context, req *models.GetAllUserRequest) (*models.	AllUserResponse, error)

}

type user struct{
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) (UserService, error){

	newService := &user{
		userRepository: repository,
	}

	return newService, nil
}

func (user user) EditUserInfo(ctx context.Context,id string, req *models.RegisterUpdateUserRequest) (*models.RegisterNewUserResponse, error){
	return nil, nil
}
func (user user) GetUserInfo(ctx context.Context, id string) (*models.UserListView, error){

	userInfo, err :=	user.userRepository.GetUser(ctx, "id", id)
	if err != nil {
		fmt.Println("Error calling userRepository.GetUser ", err)
		return nil, err
	}
	fmt.Println("Repo retrieves user info ", userInfo)
	if userInfo == nil {
		return nil, nil
	}
	return &models.UserListView{
		ID: userInfo.ID,
		User: userInfo.User,
		Role: userInfo.Role,
		State: userInfo.State,
		Register: userInfo.Register,
	}, nil
}
func (user user) 	GetAllUsersInfo(ctx context.Context, req *models.GetAllUserRequest) (*models.	AllUserResponse, error){

	users, err := user.userRepository.GetAllUser(ctx)

	if err != nil {
		fmt.Println("Error calling userRepository.GetAllUser ", err)
		return nil, err
	}
	return users, nil
}