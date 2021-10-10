package repository

import (
	"context"
	"database/sql"
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
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) (RoleRepository, error) {

		return  &roleRepository{db: db}, nil
}


func(user *roleRepository) CreateRole(ctx context.Context,data *models.Role) (*models.Role, error){

	return nil, nil
}
func(user *roleRepository) GetAllRole(ctx context.Context) ([]*models.Role, error){

	return nil, nil
}
func(user *roleRepository) GetRole(ctx context.Context, filter, value string) (*models.Role, error){

	return nil, nil
}
func(user *roleRepository) UpdateRole(ctx context.Context , data *models.Role) (*models.Role, error){

	return nil, nil
}