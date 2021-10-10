package repository


import (
	"context"
	"flypack/config"
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

}

func NewShippingRepository(repositoryType string, dbconfig *config.DBConfig) (ShippingRepository, error) {

		return  createShippingRepository(repositoryType, dbconfig)
}

func createShippingRepository(repositoryType string, dbconfig *config.DBConfig) (ShippingRepository, error){

	switch repositoryType {
	case "mysql":
		return &mysqlShippingRepository{}, nil
	default:
			return newMockShippingRepository()
	}
}

type mysqlShippingRepository struct {

}
func(user *mysqlShippingRepository) CreateShipping(ctx context.Context,data *models.Shipping) (*models.Shipping, error){

	return nil, nil
}
func(user *mysqlShippingRepository) GetAllShipping(ctx context.Context) ([]*models.Shipping, error){

	return nil, nil
}
func(user *mysqlShippingRepository) GetShipping(ctx context.Context, filter, value string) (*models.Shipping, error){

	return nil, nil
}
func(user *mysqlShippingRepository) UpdateShipping(ctx context.Context , data *models.Shipping) (*models.Shipping, error){

	return nil, nil
}
type shippingMockRepository struct {}
func newMockShippingRepository() (ShippingRepository, error){

	return &shippingMockRepository{}, nil
}
func(user *shippingMockRepository) CreateShipping(ctx context.Context,data *models.Shipping) (*models.Shipping, error){

	return nil, nil
}
func(user *shippingMockRepository) GetAllShipping(ctx context.Context) ([]*models.Shipping, error){

	return nil, nil
}
func(user *shippingMockRepository) GetShipping(ctx context.Context, filter, value string) (*models.Shipping, error){

	return nil, nil
}
func(user *shippingMockRepository) UpdateShipping(ctx context.Context , data *models.Shipping) (*models.Shipping, error){

	return nil, nil
}