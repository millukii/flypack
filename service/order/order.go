package order

import (
	"context"
	"flypack/models"
	"flypack/repository"

)

type OrderService interface {
	EditOrderInfo(ctx context.Context, req *models.RegisterNewOrderRequest) (*models.RegisterNewOrderResponse, error)
	CreateOrderInfo(ctx context.Context, req *models.RegisterNewOrderRequest) (*models.RegisterNewOrderResponse, error)
	GetOrderInfo(ctx context.Context, id string) (*models.OrderListView, error)
	GetAllOrdersInfo(ctx context.Context, req *models.GetAllOrderRequest) (*models.AllOrderResponse, error)
}

type order struct{
	orderRepository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) (OrderService){

	newService := &order{
		orderRepository: repository,
	}

	return newService
}

func (order order) EditOrderInfo(ctx context.Context, req *models.RegisterNewOrderRequest) (*models.RegisterNewOrderResponse, error){
	return nil, nil
}

func (order order) CreateOrderInfo(ctx context.Context, req *models.RegisterNewOrderRequest) (*models.RegisterNewOrderResponse, error){
	return nil, nil
}
func (order order) GetOrderInfo(ctx context.Context, id string) (*models.OrderListView, error){
	return nil, nil
}
func (order order) 	GetAllOrdersInfo(ctx context.Context, req *models.GetAllOrderRequest) (*models.	AllOrderResponse, error){
	return nil, nil
}