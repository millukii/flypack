package models

type Traveler struct {
	ID                   string `json:"id"`
	FullName             string `json:"fullname"`
	TravelerId           string `json:"travelerId"`
	Phone                string `json:"phone"`
	Address              string `json:"address"`
	Email                string `json:"email"`
	DocumentType         string `json:"documentType"`
	IdentificationNumber string `json:"identificationNumber"`
	Nacionality          string `json:"nacionality"`
	ResidenceType        string `json:"residenceType"`
	Photo                string `json:"photo"`
}

type RegisterNewTravelerRequest struct{}

type RegisterNewTravelerResponse struct{}

type TravelerListView struct{}

type GetAllTravelerRequest struct{}

type AllTravelerResponse struct{}