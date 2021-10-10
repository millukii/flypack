package repository

import (
	"context"
	"database/sql"
	"flypack/models"
)

type PeopleRepository interface {
	GetPeople(ctx context.Context, filter, value string) (*models.People, error)
	GetAllPeople(ctx context.Context) ([]*models.People, error)
	CreatePeople(ctx context.Context, 	user *models.People) (*models.People, error)
	UpdatePeople(
		ctx context.Context,
		user *models.People)(*models.People, error)
	
}

type peopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(	db *sql.DB) (PeopleRepository, error) {

		return  &peopleRepository{db: db}, nil
}

func(user *peopleRepository) CreatePeople(ctx context.Context,data *models.People) (*models.People, error){

	return nil, nil
}
func(user *peopleRepository) GetAllPeople(ctx context.Context) ([]*models.People, error){

	return nil, nil
}
func(user *peopleRepository) GetPeople(ctx context.Context, filter, value string) (*models.People, error){

	return nil, nil
}
func(user *peopleRepository) UpdatePeople(ctx context.Context , data *models.People) (*models.People, error){

	return nil, nil
}