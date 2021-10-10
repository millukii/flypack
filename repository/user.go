package repository

import (
	"context"
	"database/sql"
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
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (UserRepository, error) {

		return &userRepository{db: db}, nil
}

func(user *userRepository) CreateUser(ctx context.Context,data *models.User) (*models.User, error){

	return nil, nil
}
func(user *userRepository) GetAllUser(ctx context.Context) ([]*models.User, error){

	return nil, nil
}
func(user *userRepository) GetUser(ctx context.Context, filter, value string) (*models.User, error){

	return nil, nil
}
func(user *userRepository) UpdateUser(ctx context.Context , data *models.User) (*models.User, error){

	return nil, nil
}