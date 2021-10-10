package repository


import (
	"context"
	"flypack/config"
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

}

func NewPeopleRepository(repositoryType string, dbconfig *config.DBConfig) (PeopleRepository, error) {

		return  createPeopleRepository(repositoryType, dbconfig)
}

func createPeopleRepository(repositoryType string, dbconfig *config.DBConfig) (PeopleRepository, error){

	switch repositoryType {
	case "mysql":
		return &mysqlPeopleRepository{}, nil
	default:
			return newMockPeopleRepository()
	}
}

type mysqlPeopleRepository struct {

}
func(user *mysqlPeopleRepository) CreatePeople(ctx context.Context,data *models.People) (*models.People, error){

	return nil, nil
}
func(user *mysqlPeopleRepository) GetAllPeople(ctx context.Context) ([]*models.People, error){

	return nil, nil
}
func(user *mysqlPeopleRepository) GetPeople(ctx context.Context, filter, value string) (*models.People, error){

	return nil, nil
}
func(user *mysqlPeopleRepository) UpdatePeople(ctx context.Context , data *models.People) (*models.People, error){

	return nil, nil
}
type peopleMockRepository struct {}
func newMockPeopleRepository() (PeopleRepository, error){

	return &peopleMockRepository{}, nil
}
func(user *peopleMockRepository) CreatePeople(ctx context.Context,data *models.People) (*models.People, error){

	return nil, nil
}
func(user *peopleMockRepository) GetAllPeople(ctx context.Context) ([]*models.People, error){

	return nil, nil
}
func(user *peopleMockRepository) GetPeople(ctx context.Context, filter, value string) (*models.People, error){

	return nil, nil
}
func(user *peopleMockRepository) UpdatePeople(ctx context.Context , data *models.People) (*models.People, error){

	return nil, nil
}