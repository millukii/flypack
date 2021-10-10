package shipping

import (
	"context"
	"flypack/repository"

)

type ShippingStateService interface {
	AddShippingState (ctx context.Context, name string) (string, error)
  EditShippingState (ctx context.Context, name string) (string, error)
	GetShippingStates (ctx context.Context ) ([]string, error)
	GetShippingState (ctx context.Context,id string ) (string, error)
	DeactivateShippingState (ctx context.Context, name string) (string, error)
}

type shippingStateService struct{
	repository repository.ShippingStateRepository
}

func NewShippingStateService(repository repository.ShippingStateRepository) (ShippingStateService, error) {

	return &shippingStateService{repository: repository}, nil
}

	func (r shippingStateService) AddShippingState (ctx context.Context, name string) (string, error){
		return "", nil
	}
 	func (r shippingStateService) EditShippingState (ctx context.Context, name string) (string, error){
		 		return "", nil
	 }
	func (r shippingStateService)	DeactivateShippingState (ctx context.Context, name string) (string, error){
				return "", nil
	}

	func (r shippingStateService)		GetShippingStates (ctx context.Context ) ([]string, error){

		return []string{}, nil
	}
		func (r shippingStateService)		GetShippingState (ctx context.Context,id string ) (string, error){

		return "", nil
	}