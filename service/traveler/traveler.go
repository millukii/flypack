package traveler

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type TravelerService interface {
	EditTravelerInfo(ctx context.Context, req *models.RegisterNewTravelerRequest) (*models.RegisterNewTravelerResponse, error)
	CreateTravelerInfo(ctx context.Context, req *models.RegisterNewTravelerRequest) (*models.RegisterNewTravelerResponse, error)
	GetTravelerInfo(ctx context.Context, id string) (*models.TravelerListView, error)
	GetAllTravelersInfo(ctx context.Context, req *models.GetAllTravelerRequest) (*models.AllTravelerResponse, error)
}

type traveler struct{
	travelerRepository repository.TravelerRepository
}

func NewTravelerService(repository repository.TravelerRepository) (TravelerService){

	newService := &traveler{
		travelerRepository: repository,
	}

	return newService
}

func (traveler traveler) EditTravelerInfo(ctx context.Context, req *models.RegisterNewTravelerRequest) (*models.RegisterNewTravelerResponse, error){
	return nil, nil
}

func (traveler traveler) CreateTravelerInfo(ctx context.Context, req *models.RegisterNewTravelerRequest) (*models.RegisterNewTravelerResponse, error){
	return nil, nil
}
func (traveler traveler) GetTravelerInfo(ctx context.Context, id string) (*models.TravelerListView, error){
	return nil, nil
}
func (traveler traveler) 	GetAllTravelersInfo(ctx context.Context, req *models.GetAllTravelerRequest) (*models.	AllTravelerResponse, error){
	return nil, nil
}