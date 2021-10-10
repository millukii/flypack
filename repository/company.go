package repository

import (
	"context"
	"database/sql"
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
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) (CompanyRepository, error) {

		return &companyRepository{db: db}, nil
}

func(user *companyRepository) CreateCompany(ctx context.Context,data *models.Company) (*models.Company, error){

	return nil, nil
}
func(user *companyRepository) GetAllCompany(ctx context.Context) ([]*models.Company, error){

	return nil, nil
}
func(user *companyRepository) GetCompany(ctx context.Context, filter, value string) (*models.Company, error){

	return nil, nil
}
func(user *companyRepository) UpdateCompany(ctx context.Context , data *models.Company) (*models.Company, error){

	return nil, nil
}