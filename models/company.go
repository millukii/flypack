package models

type Company struct {
	ID          string `json:"id"`
	Rut         int    `json:"rut"`
	DV          string `json:"dv"`
	Razon       string `json:"razon"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Commune     string `json:"commune"`
	LegalPerson int    `json:"people_id"`
}
type RegisterNewCompanyRequest struct {}
type RegisterNewCompanyResponse struct {}
type GetAllCompanyRequest struct {}
type CompanyListView struct {}
type AllCompanyResponse struct {}