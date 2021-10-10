package shipping

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type ShippingService interface {
	GetShippingById(ctx context.Context, id string) (*models.ShippingListView, error)
	CreateShipping(ctx context.Context, req *models.CreateShippingRequest) (*models.ShippingListView, error)
	 GetAllShippingInfo(ctx context.Context, id string) (*models.ShippingListView, error)
	UpdateShippingInfo(ctx context.Context, id string) (*models.ShippingListView, error)
	CancelShipping(ctx context.Context, id string) (*models.ShippingListView, error)
}

type shippingService struct{
	repository *repository.ShippingRepository
}

func NewShippingService(repository *repository.ShippingRepository) (ShippingService, error) {

	return &shippingService{
		repository: repository,
	}, nil
}

func (svc shippingService)		GetShippingById(ctx context.Context, id string) (*models.ShippingListView, error){

	return nil, nil
}

func (svc shippingService)		CreateShipping(ctx context.Context, req *models.CreateShippingRequest) (*models.ShippingListView, error){

	return nil, nil
}

func (svc shippingService)	 GetAllShippingInfo(ctx context.Context, id string) (*models.ShippingListView, error){
		return nil, nil
}


func (svc shippingService)	UpdateShippingInfo(ctx context.Context, id string) (*models.ShippingListView, error){
		return nil, nil
}
func (svc shippingService)	CancelShipping(ctx context.Context, id string) (*models.ShippingListView, error){
		return nil, nil
}