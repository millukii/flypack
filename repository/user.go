package repository

import (
	"context"
	"flypack/config"
	"flypack/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, filter, value string) (*models.User, error)
	GetAllUser(ctx context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, 	user *models.User) (*models.User, error)
	UpdateUser(
		ctx context.Context,
		user *models.User)(*models.User, error)
	
}

type userRepository struct {

}

func NewUserRepository(repositoryType string, dbconfig *config.DBConfig) (UserRepository, error) {

		return  createRepo(repositoryType, dbconfig)
}

func createRepo(repositoryType string, dbconfig *config.DBConfig) (UserRepository, error){

	switch repositoryType {
	case "mysql":
		return &mysqlUserRepository{}, nil
	default:
			return newMockUserRepository()
	}
}

type mysqlUserRepository struct {

}
func(user *mysqlUserRepository) CreateUser(ctx context.Context,data *models.User) (*models.User, error){

	return nil, nil
}
func(user *mysqlUserRepository) GetAllUser(ctx context.Context) ([]*models.User, error){

	return nil, nil
}
func(user *mysqlUserRepository) GetUser(ctx context.Context, filter, value string) (*models.User, error){

	return nil, nil
}
func(user *mysqlUserRepository) UpdateUser(ctx context.Context , data *models.User) (*models.User, error){

	return nil, nil
}
type userMockRepository struct {}
func newMockUserRepository() (UserRepository, error){

	return &userMockRepository{}, nil
}
func(user *userMockRepository) CreateUser(ctx context.Context,data *models.User) (*models.User, error){

	return nil, nil
}
func(user *userMockRepository) GetAllUser(ctx context.Context) ([]*models.User, error){

	return nil, nil
}
func(user *userMockRepository) GetUser(ctx context.Context, filter, value string) (*models.User, error){

	return nil, nil
}
func(user *userMockRepository) UpdateUser(ctx context.Context , data *models.User) (*models.User, error){

	return nil, nil
}