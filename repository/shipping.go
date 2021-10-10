package repository

import (
	"context"
	"database/sql"
	"flypack/models"
)

type ShippingRepository interface {
	GetShipping(ctx context.Context, filter, value string) (*models.Shipping, error)
	GetAllShipping(ctx context.Context) ([]*models.Shipping, error)
	CreateShipping(ctx context.Context, 	user *models.Shipping) (*models.Shipping, error)
	UpdateShipping(
		ctx context.Context,
		user *models.Shipping)(*models.Shipping, error)
	
}

type shippingRepository struct {
	db *sql.DB
}

func NewShippingRepository(	db *sql.DB) (ShippingRepository, error) {

		return  &shippingRepository{db: db}, nil
}


func(user *shippingRepository) CreateShipping(ctx context.Context,data *models.Shipping) (*models.Shipping, error){

	return nil, nil
}
func(user *shippingRepository) GetAllShipping(ctx context.Context) ([]*models.Shipping, error){

	return nil, nil
}
func(user *shippingRepository) GetShipping(ctx context.Context, filter, value string) (*models.Shipping, error){

	return nil, nil
}
func(user *shippingRepository) UpdateShipping(ctx context.Context , data *models.Shipping) (*models.Shipping, error){

	return nil, nil
}