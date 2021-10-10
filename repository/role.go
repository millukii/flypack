package repository

import (
	"context"
	"flypack/config"
	"flypack/models"

)

type RoleRepository interface {
	GetRole(ctx context.Context, filter, value string) (*models.Role, error)
	GetAllRole(ctx context.Context) ([]*models.Role, error)
	CreateRole(ctx context.Context, 	user *models.Role) (*models.Role, error)
	UpdateRole(
		ctx context.Context,
		user *models.Role)(*models.Role, error)
	
}

type roleRepository struct {

}

func NewRoleRepository(repositoryType string, dbconfig *config.DBConfig) (RoleRepository, error) {

		return  createRoleRepository(repositoryType, dbconfig)
}

func createRoleRepository(repositoryType string, dbconfig *config.DBConfig) (RoleRepository, error){

	switch repositoryType {
	case "mysql":
		return &mysqlRoleRepository{}, nil
	default:
			return newMockRoleRepository()
	}
}

type mysqlRoleRepository struct {

}
func(user *mysqlRoleRepository) CreateRole(ctx context.Context,data *models.Role) (*models.Role, error){

	return nil, nil
}
func(user *mysqlRoleRepository) GetAllRole(ctx context.Context) ([]*models.Role, error){

	return nil, nil
}
func(user *mysqlRoleRepository) GetRole(ctx context.Context, filter, value string) (*models.Role, error){

	return nil, nil
}
func(user *mysqlRoleRepository) UpdateRole(ctx context.Context , data *models.Role) (*models.Role, error){

	return nil, nil
}
type roleMockRepository struct {}
func newMockRoleRepository() (RoleRepository, error){

	return &roleMockRepository{}, nil
}
func(user *roleMockRepository) CreateRole(ctx context.Context,data *models.Role) (*models.Role, error){

	return nil, nil
}
func(user *roleMockRepository) GetAllRole(ctx context.Context) ([]*models.Role, error){

	return nil, nil
}
func(user *roleMockRepository) GetRole(ctx context.Context, filter, value string) (*models.Role, error){

	return nil, nil
}
func(user *roleMockRepository) UpdateRole(ctx context.Context , data *models.Role) (*models.Role, error){

	return nil, nil
}