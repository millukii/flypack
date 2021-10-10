package user

import (
	"context"
	"flypack/repository"

)

type RoleService interface {
	AddRole (ctx context.Context, name string) (string, error)
  EditRole (ctx context.Context, name string) (string, error)
	GetRoles (ctx context.Context ) ([]string, error)
	GetRole (ctx context.Context,id string ) (string, error)
	DeactivateRole (ctx context.Context, name string) (string, error)
}

type roleService struct{
	repository repository.RoleRepository
}

func NewRoleService(repository repository.RoleRepository) (RoleService, error) {

	return &roleService{repository: repository}, nil
}

	func (r roleService) AddRole (ctx context.Context, name string) (string, error){
		return "", nil
	}
 	func (r roleService) EditRole (ctx context.Context, name string) (string, error){
		 		return "", nil
	 }
	func (r roleService)	DeactivateRole (ctx context.Context, name string) (string, error){
				return "", nil
	}

	func (r roleService)		GetRoles (ctx context.Context ) ([]string, error){

		return []string{}, nil
	}
		func (r roleService)		GetRole (ctx context.Context,id string ) (string, error){

		return "", nil
	}