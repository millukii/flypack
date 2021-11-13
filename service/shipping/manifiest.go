package shipping

import (
	"context"
	"flypack/models"
	"flypack/repository"
)
type ManifiestService interface {
	EditManifiestInfo(ctx context.Context, req *models.RegisterNewManifiestRequest) (*models.RegisterNewManifiestResponse, error)
	CreateManifiestInfo(ctx context.Context, req *models.RegisterNewManifiestRequest) (*models.RegisterNewManifiestResponse, error)
	GetManifiestInfo(ctx context.Context, id string) (*models.ManifiestListView, error)
	GetAllManifiestsInfo(ctx context.Context, req *models.GetAllManifiestRequest) (*models.AllManifiestResponse, error)
}

type manifiest struct {
	manifiestRepository repository.ManifiestRepository
}

func NewManifiestService(repository repository.ManifiestRepository) ManifiestService {

	newService := &manifiest{
		manifiestRepository: repository,
	}

	return newService
}

func (manifiest manifiest) EditManifiestInfo(ctx context.Context, req *models.RegisterNewManifiestRequest) (*models.RegisterNewManifiestResponse, error) {
	return nil, nil
}

func (manifiest manifiest) CreateManifiestInfo(ctx context.Context, req *models.RegisterNewManifiestRequest) (*models.RegisterNewManifiestResponse, error) {
	return nil, nil
}
func (manifiest manifiest) GetManifiestInfo(ctx context.Context, id string) (*models.ManifiestListView, error) {
	return nil, nil
}
func (manifiest manifiest) GetAllManifiestsInfo(ctx context.Context, req *models.GetAllManifiestRequest) (*models.AllManifiestResponse, error) {
	return nil, nil
}