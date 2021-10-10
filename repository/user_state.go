package repository

import (
	"context"
	"flypack/config"
	"flypack/models"

)

type UserStateRepository interface {
	GetUserState(ctx context.Context, filter, value string) (*models.UserState, error)
	GetAllUserState(ctx context.Context) ([]*models.UserState, error)
	CreateUserState(ctx context.Context, 	user *models.UserState) (*models.UserState, error)
	UpdateUserState(
		ctx context.Context,
		user *models.UserState)(*models.UserState, error)
	
}

type userStateRepository struct {

}

func NewUserStateRepository(repositoryType string, dbconfig *config.DBConfig) (UserStateRepository, error) {

		return  createUserStateRepository(repositoryType, dbconfig)
}

func createUserStateRepository(repositoryType string, dbconfig *config.DBConfig) (UserStateRepository, error){

	switch repositoryType {
	case "mysql":
		return &mysqlUserStateRepository{}, nil
	default:
			return newMockUserStateRepository()
	}
}

type mysqlUserStateRepository struct {

}
func(user *mysqlUserStateRepository) CreateUserState(ctx context.Context,data *models.UserState) (*models.UserState, error){

	return nil, nil
}
func(user *mysqlUserStateRepository) GetAllUserState(ctx context.Context) ([]*models.UserState, error){

	return nil, nil
}
func(user *mysqlUserStateRepository) GetUserState(ctx context.Context, filter, value string) (*models.UserState, error){

	return nil, nil
}
func(user *mysqlUserStateRepository) UpdateUserState(ctx context.Context , data *models.UserState) (*models.UserState, error){

	return nil, nil
}
type userStateMockRepository struct {}
func newMockUserStateRepository() (UserStateRepository, error){

	return &userStateMockRepository{}, nil
}
func(user *userStateMockRepository) CreateUserState(ctx context.Context,data *models.UserState) (*models.UserState, error){

	return nil, nil
}
func(user *userStateMockRepository) GetAllUserState(ctx context.Context) ([]*models.UserState, error){

	return nil, nil
}
func(user *userStateMockRepository) GetUserState(ctx context.Context, filter, value string) (*models.UserState, error){

	return nil, nil
}
func(user *userStateMockRepository) UpdateUserState(ctx context.Context , data *models.UserState) (*models.UserState, error){

	return nil, nil
}