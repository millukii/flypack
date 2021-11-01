package user

import (
	"context"
	"flypack/errors"
	"flypack/logger"
	"flypack/models"
	"flypack/repository"
	"fmt"

)

type UserService interface {
	EditUserInfo(ctx context.Context,id string, req *models.RegisterUpdateUserRequest) (*models.UserListView, error)
	GetUserInfo(ctx context.Context, id string) (*models.UserListView, error)
	GetAllUsersInfo(ctx context.Context, req *models.GetAllUserRequest) (*models.	AllUserResponse, error)

}

type userService struct{
	userRepository repository.UserRepository
	logger logger.ApplicationLogger
}

func NewUserService(repository repository.UserRepository, logger logger.ApplicationLogger) (UserService){

	newService := &userService{
		userRepository: repository,
		logger: logger,
	}

	return newService
}

func (userSvc userService) EditUserInfo(ctx context.Context,id string, req *models.RegisterUpdateUserRequest) (*models.UserListView, error){
	userSvc.logger.InitFunction()
	defer userSvc.logger.EndFunction()

	updatedUser, err :=	userSvc.userRepository.UpdateUser(ctx,&models.User{
		User: req.User, 
		Role: req.Role,
		State: req.State,
	})
	if err != nil {
		userSvc.logger.Error(err, errors.ErrRepositoryUpdateRecord.Error())
		return nil, err
	}
	return &models.UserListView{
		User: updatedUser.User,
		Role: updatedUser.Role,
		State: updatedUser.State,
	}, nil
}
func (userSvc userService) GetUserInfo(ctx context.Context, id string) (*models.UserListView, error){

	userSvc.logger.InitFunction()
	defer userSvc.logger.EndFunction()
	
	userInfo, err :=	userSvc.userRepository.GetUser(ctx, "id", id)

	if err != nil {
		userSvc.logger.Error(err, errors.ErrRepositoryGetRecord.Error())
		return nil, err
	}
	userSvc.logger.Info(fmt.Sprintf("Repo retrieves user info %+v", userInfo))
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
func (userSvc userService) 	GetAllUsersInfo(ctx context.Context, req *models.GetAllUserRequest) (*models.	AllUserResponse, error){
	userSvc.logger.InitFunction()
	defer userSvc.logger.EndFunction()

	users, err := userSvc.userRepository.GetAllUser(ctx)

	if err != nil {
		userSvc.logger.Error(err, errors.ErrRepositoryGetAll.Error())
		return nil, err
	}
	return users, nil
}