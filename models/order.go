package models

type Order struct {
	ID              string        `json:"id"`
	PurchaseOrderID string        `json:"purchaseOrderId"`
	CustomerOrderID string        `json:"customerOrderId"`
	CustomerEmailID string        `json:"customerEmailId"`
	OrderType       string        `json:"orderType"`
	OrderDate       string        `json:"orderDate"`
	OrderStatus     string        `json:"orderStatus"`
	ShippingInfo    *ShippingInfo `json:"shippingInfo"`
	OrderLines      []OrderLine   `json:"orderLines"`
}
type OrderLine struct {
	LineNumber int         `json:"lineNumber"`
	Item       *Item       `json:"item"`
	Charges    []*Charge   `json:"charges"`
	Taxes      []*Tax      `json:"taxes"`
	Discounts  []*Discount `json:"discounts"`
}
type Charge struct {
	ChargeType string `json:"chargeType"`
	ChargeName string `json:"chargeName"`
	Currency   string `json:"currency"`
	Amount     int    `json:"amount"`
}
type Tax struct {
	TaxName  string  `json:"taxName"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}
type Discount struct {
	DiscountName string  `json:"discountName"`
	Currency     string  `json:"currency"`
	Amount       float64 `json:"amount"`
}
type Item struct {
	ProductName string `json:"productName"`
	Sku         string `json:"sku"`
}
type ShippingInfo struct {
	Phone                 string         `json:"phone"`
	EstimatedDeliveryDate int64          `json:"estimatedDeliveryDate"`
	EstimatedShipDate     int64          `json:"estimatedShipDate"`
	MethodCode            string         `json:"methodCode"`
	PostalAddress         *PostalAddress `json:"postalAddress"`
}
type PostalAddress struct {
	Name        string `json:"name"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
	Country     string `json:"country"`
	AddressType string `json:"addressType"`
}

type GetAllOrderRequest struct {
}
type RegisterNewOrderRequest struct {
}

type RegisterNewOrderResponse struct{}

type OrderListView struct{
	ID              string        `json:"id"`
	PurchaseOrderID string        `json:"purchaseOrderId"`
	CustomerOrderID string        `json:"customerOrderId"`
	CustomerEmailID string        `json:"customerEmailId"`
	OrderType       string        `json:"orderType"`
	OrderDate       string        `json:"orderDate"`
	OrderStatus     string        `json:"orderStatus"`
}

type AllOrderResponse struct{
	Orders []OrderListView `json:"orders"`
	Count int `json:"count"`
}