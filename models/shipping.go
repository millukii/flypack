package models

type Shipping struct {
	ID            string  `json:"id"`
	OrderNumber   int     `json:"order_nro"`
	TickerNumber  string  `json:"boleta_nro"`
	ShippingType  string  `json:"type"`
	Value         float64 `json:"value"`
	Date          string  `json:"date"`
	ShippingState int     `json:"shipping_states_id"`
	Company       int     `json:"companies_id"`
	Delivery      int     `json:"delivery_id"`
	Client        int     `json:"client_id"`
}

type CreateShippingRequest struct {
	OrderNumber  int    `json:"order_nro"`
	ShippingType string `json:"type"`
	Company      int    `json:"companies_id"`
	Delivery     int    `json:"delivery_id"`
	Client       int    `json:"client_id"`
}

type ShippingListView struct {
	OrderNumber   int     `json:"order_nro"`
	TickerNumber  string  `json:"boleta_nro"`
	ShippingType  string  `json:"type"`
	Value         float64 `json:"value"`
	Date          string  `json:"date"`
	ShippingState int     `json:"shipping_states_id"`
	Company       int     `json:"companies_id"`
	Delivery      int     `json:"delivery_id"`
	Client        int     `json:"client_id"`
}

type AllShippingResponse struct {
	ShippingListView []*ShippingListView `json:"shippingListView"`
	Count            int                 `json:"count"`
}

type RegisterNewShippingResponse struct {}
type GetAllShippingRequest struct {}