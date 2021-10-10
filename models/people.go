package models

type People struct {
	ID       string `json:"id"`
	Rut      int    `json:"rut"`
	DV       string `json:"dv"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Address  string `json:"address"`
	City     string `json:"city"`
	Commune  string `json:"commune"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Profile  int    `json:"profile_id"`
}

type RegisterNewPeopleRequest struct {}
type RegisterNewPeopleResponse struct {}
type GetAllPeopleRequest struct {}
type PeopleListView struct {}
type AllPeopleResponse struct {}