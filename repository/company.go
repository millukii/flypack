package repository

import (
	"context"
	"flypack/config"
	"flypack/models"
)

type CompanyRepository interface {
	GetCompany(ctx context.Context, filter, value string) (*models.Company, error)
	GetAllCompany(ctx context.Context) ([]*models.Company, error)
	CreateCompany(ctx context.Context, 	user *models.Company) (*models.Company, error)
	UpdateCompany(
		ctx context.Context,
		user *models.Company)(*models.Company, error)
	
}

type companyRepository struct {

}

func NewCompanyRepository(repositoryType string, dbconfig *config.DBConfig) (CompanyRepository, error) {

		return  createCompanyRepository(repositoryType, dbconfig)
}

func createCompanyRepository(repositoryType string, dbconfig *config.DBConfig) (CompanyRepository, error){

	switch repositoryType {
	case "mysql":
		return &mysqlCompanyRepository{}, nil
	default:
			return newMockCompanyRepository()
	}
}

type mysqlCompanyRepository struct {

}
func(user *mysqlCompanyRepository) CreateCompany(ctx context.Context,data *models.Company) (*models.Company, error){

	return nil, nil
}
func(user *mysqlCompanyRepository) GetAllCompany(ctx context.Context) ([]*models.Company, error){

	return nil, nil
}
func(user *mysqlCompanyRepository) GetCompany(ctx context.Context, filter, value string) (*models.Company, error){

	return nil, nil
}
func(user *mysqlCompanyRepository) UpdateCompany(ctx context.Context , data *models.Company) (*models.Company, error){

	return nil, nil
}
type companyMockRepository struct {}
func newMockCompanyRepository() (CompanyRepository, error){

	return &companyMockRepository{}, nil
}
func(user *companyMockRepository) CreateCompany(ctx context.Context,data *models.Company) (*models.Company, error){

	return nil, nil
}
func(user *companyMockRepository) GetAllCompany(ctx context.Context) ([]*models.Company, error){

	return nil, nil
}
func(user *companyMockRepository) GetCompany(ctx context.Context, filter, value string) (*models.Company, error){

	return nil, nil
}
func(user *companyMockRepository) UpdateCompany(ctx context.Context , data *models.Company) (*models.Company, error){

	return nil, nil
}