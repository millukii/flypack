package models

type ShipmentPackage struct {
	Id          string   `json:"id"`
	Size        string   `json:"size"`
	OrderNumber string   `json:"orderNumber"`
	Items       []string `json:"items"`
}

type RegisterNewShipmentPackagesRequest struct{}
type RegisterNewShipmentPackagesResponse struct{}
type ShipmentPackagesListView struct{}
type GetAllShipmentPackagesRequest struct{}
type AllShipmentPackagesResponse struct {
	PackagesList []*ShipmentPackage `json:"shipmentPackages"`
	Count        int                `json:"count"`
}