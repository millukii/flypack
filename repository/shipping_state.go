package repository

import (
	"context"
	"database/sql"
	"flypack/models"
)

type ShippingStateRepository interface {
	GetShippingState(ctx context.Context, filter, value string) (*models.ShippingState, error)
	GetAllShippingState(ctx context.Context) ([]*models.ShippingState, error)
	CreateShippingState(ctx context.Context, 	user *models.ShippingState) (*models.ShippingState, error)
	UpdateShippingState(
		ctx context.Context,
		user *models.ShippingState)(*models.ShippingState, error)
	
}

type shippingStateRepository struct {
	db *sql.DB
}

func NewShippingStateRepository(db *sql.DB) (ShippingStateRepository, error) {

		return  &shippingStateRepository{db: db}, nil
}


func(user *shippingStateRepository) CreateShippingState(ctx context.Context,data *models.ShippingState) (*models.ShippingState, error){

	return nil, nil
}
func(user *shippingStateRepository) GetAllShippingState(ctx context.Context) ([]*models.ShippingState, error){

	return nil, nil
}
func(user *shippingStateRepository) GetShippingState(ctx context.Context, filter, value string) (*models.ShippingState, error){

	return nil, nil
}
func(user *shippingStateRepository) UpdateShippingState(ctx context.Context , data *models.ShippingState) (*models.ShippingState, error){

	return nil, nil
}