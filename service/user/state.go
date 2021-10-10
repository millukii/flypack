package user

import (
	"context"
	"flypack/repository"
)

type UserStateService interface {
	AddUserState (ctx context.Context, name string) (string, error)
	GetUserState (ctx context.Context, id string) (string, error)
	GetUserStates (ctx context.Context) ([]string, error)
  EditUserState (ctx context.Context, name string) (string, error)
	DeactivateUserState (ctx context.Context, name string) (string, error)
}

type userStateService struct{
	repository repository.UserStateRepository
}

func NewUserStateService(repository repository.UserStateRepository) (UserStateService, error) {

	return &userStateService{repository: repository}, nil
}

	func (r userStateService) GetUserStates (ctx context.Context) ([]string, error){
		return []string{},nil 
	}
	func (r userStateService) AddUserState (ctx context.Context, name string) (string, error){
		return "", nil
	}
 	func (r userStateService) EditUserState (ctx context.Context, name string) (string, error){
		 		return "", nil
	 }
	func (r userStateService) GetUserState (ctx context.Context, id string) (string, error){
		 		return "", nil
	 }
	func (r userStateService)	DeactivateUserState (ctx context.Context, name string) (string, error){
				return "", nil
	}