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

	var userCreated *models.User
	tx, err := user.db.Begin()

	if err != nil {
		return nil, err
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO users(
			password,
			people_id, 
			rol_id, 
			user,
			user_state_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id`)

		if err != nil {
			return nil, err
		}

		defer stmt.Close()

		err = stmt.QueryRow(
			data.Password,
			data.Register,
			data.Rol,
			data.User, 
			data.State,
		).Scan(&userCreated)

		if err != nil {
			return nil, err
		}
	}

	{
		err := tx.Commit()

		if err != nil {
			return nil, err
		}
	}

	return userCreated, nil
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