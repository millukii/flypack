package people

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type PeopleService interface {
	EditPeopleInfo(ctx context.Context, req *models.RegisterNewPeopleRequest) (*models.RegisterNewPeopleResponse, error)
	RegisterPerson(ctx context.Context, req *models.RegisterNewPeopleRequest) (*models.RegisterNewPeopleResponse, error)
	GetPeopleInfo(ctx context.Context, id string) (*models.PeopleListView, error)
	GetAllPeoplesInfo(ctx context.Context, req *models.GetAllPeopleRequest) (*models.AllPeopleResponse, error)

}

type people struct{
	peopleRepository repository.PeopleRepository
}

func NewPeopleService(repository repository.PeopleRepository) (PeopleService){

	newService := &people{
		peopleRepository: repository,
	}

	return newService
}

func (people people) RegisterPerson(ctx context.Context, req *models.RegisterNewPeopleRequest) (*models.RegisterNewPeopleResponse, error){
	return nil, nil
}
func (people people) EditPeopleInfo(ctx context.Context, req *models.RegisterNewPeopleRequest) (*models.RegisterNewPeopleResponse, error){
	return nil, nil
}
func (people people) GetPeopleInfo(ctx context.Context, id string) (*models.PeopleListView, error){
	return nil, nil
}
func (people people) 	GetAllPeoplesInfo(ctx context.Context, req *models.GetAllPeopleRequest) (*models.	AllPeopleResponse, error){
	return nil, nil
}