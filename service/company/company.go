package company

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type CompanyService interface {
	EditCompanyInfo(ctx context.Context, req *models.RegisterNewCompanyRequest) (*models.RegisterNewCompanyResponse, error)
	GetCompanyInfo(ctx context.Context, id string) (*models.CompanyListView, error)
	GetAllCompanysInfo(ctx context.Context, req *models.GetAllCompanyRequest) (*models.	AllCompanyResponse, error)

}

type company struct{
	companyRepository *repository.CompanyRepository
}

func NewCompanyService(repository *repository.CompanyRepository) (CompanyService, error){

	newService := &company{
		companyRepository: repository,
	}

	return newService, nil
}

func (company company) EditCompanyInfo(ctx context.Context, req *models.RegisterNewCompanyRequest) (*models.RegisterNewCompanyResponse, error){
	return nil, nil
}
func (company company) GetCompanyInfo(ctx context.Context, id string) (*models.CompanyListView, error){
	return nil, nil
}
func (company company) 	GetAllCompanysInfo(ctx context.Context, req *models.GetAllCompanyRequest) (*models.	AllCompanyResponse, error){
	return nil, nil
}