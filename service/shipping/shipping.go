package shipping

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type ShippingService interface {
	GetShippingById(ctx context.Context, id string) (*models.ShippingListView, error)
	CreateShipping(ctx context.Context, req *models.CreateShippingRequest) (*models.ShippingListView, error)
	 GetAllShippingInfo(ctx context.Context,req *models.GetAllShippingRequest) ([]*models.Shipping, error)
	UpdateShippingInfo(ctx context.Context, id string) (*models.ShippingListView, error)
	CancelShipping(ctx context.Context, id string) (*models.ShippingListView, error)
}

type shippingService struct{
	repository repository.ShippingRepository
}

func NewShippingService(repository repository.ShippingRepository) (ShippingService) {

	return &shippingService{
		repository: repository,
	}
}

func (svc shippingService)		GetShippingById(ctx context.Context, id string) (*models.ShippingListView, error){

	return nil, nil
}

func (svc shippingService)		CreateShipping(ctx context.Context, req *models.CreateShippingRequest) (*models.ShippingListView, error){

	return nil, nil
}

func (svc shippingService)	 GetAllShippingInfo(ctx context.Context,req *models.GetAllShippingRequest) ([]*models.Shipping, error){
		
	shippings, err :=	svc.repository.GetAllShipping(ctx)
	if err != nil {
		return nil, err
	}
		return shippings, nil
}


func (svc shippingService)	UpdateShippingInfo(ctx context.Context, id string) (*models.ShippingListView, error){
		return nil, nil
}
func (svc shippingService)	CancelShipping(ctx context.Context, id string) (*models.ShippingListView, error){
		return nil, nil
}