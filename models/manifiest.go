package models

type Manifiest struct {
	ID                     string             `json:"id"`
	ManifiestNumber        string             `json:"manifiestNumber"`
	DeliveryDate           string             `json:"deliveryDate"`
	City                   string             `json:"city"`
	Commune                string             `json:"commune"`
	SmallPackagesCapacity  int                `json:"smallPackagesCapacity"`
	MediumPackagesCapacity int                `json:"mediumPackagesCapacity"`
	XtraPackagesCapacity   int                `json:"xtraPackagesCapacity"`
	PackageList            []*ShipmentPackage `json:"shipmentPackages"`
}

type RegisterNewManifiestRequest struct{}

type ManifiestListView struct{}
type GetAllManifiestRequest struct{}

type RegisterNewManifiestResponse struct {
}
type AllManifiestResponse struct {
	Manifiests []*Manifiest `json:"manifiest"`
	Count      int          `json:"count"`
}