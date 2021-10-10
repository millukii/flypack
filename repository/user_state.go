package repository

import (
	"context"
	"database/sql"
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
 db *sql.DB
}

func NewUserStateRepository( db *sql.DB) (UserStateRepository, error) {

		return  &userStateRepository{db: db}, nil
}


func(user *userStateRepository) CreateUserState(ctx context.Context,data *models.UserState) (*models.UserState, error){

	return nil, nil
}
func(user *userStateRepository) GetAllUserState(ctx context.Context) ([]*models.UserState, error){

	return nil, nil
}
func(user *userStateRepository) GetUserState(ctx context.Context, filter, value string) (*models.UserState, error){

	return nil, nil
}
func(user *userStateRepository) UpdateUserState(ctx context.Context , data *models.UserState) (*models.UserState, error){

	return nil, nil
}