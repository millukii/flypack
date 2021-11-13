package shipping

import (
	"context"
	"flypack/models"
	"flypack/repository"
)
type ShipmentPackagesService interface {
	EditShipmentPackagesInfo(ctx context.Context, req *models.RegisterNewShipmentPackagesRequest) (*models.RegisterNewShipmentPackagesResponse, error)
	CreateShipmentPackagesInfo(ctx context.Context, req *models.RegisterNewShipmentPackagesRequest) (*models.RegisterNewShipmentPackagesResponse, error)
	GetShipmentPackagesInfo(ctx context.Context, id string) (*models.ShipmentPackagesListView, error)
	GetAllShipmentPackagessInfo(ctx context.Context, req *models.GetAllShipmentPackagesRequest) (*models.AllShipmentPackagesResponse, error)
}

type shipmentPackages struct {
	shipmentPackagesRepository repository.ShipmentPackagesRepository
}

func NewShipmentPackagesService(repository repository.ShipmentPackagesRepository) ShipmentPackagesService {

	newService := &shipmentPackages{
		shipmentPackagesRepository: repository,
	}

	return newService
}

func (shipmentPackages shipmentPackages) EditShipmentPackagesInfo(ctx context.Context, req *models.RegisterNewShipmentPackagesRequest) (*models.RegisterNewShipmentPackagesResponse, error) {
	return nil, nil
}

func (shipmentPackages shipmentPackages) CreateShipmentPackagesInfo(ctx context.Context, req *models.RegisterNewShipmentPackagesRequest) (*models.RegisterNewShipmentPackagesResponse, error) {
	return nil, nil
}
func (shipmentPackages shipmentPackages) GetShipmentPackagesInfo(ctx context.Context, id string) (*models.ShipmentPackagesListView, error) {
	return nil, nil
}
func (shipmentPackages shipmentPackages) GetAllShipmentPackagessInfo(ctx context.Context, req *models.GetAllShipmentPackagesRequest) (*models.AllShipmentPackagesResponse, error) {
	return nil, nil
}