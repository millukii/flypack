package handlers

import (
	"flypack/handlers/mocks"
	"flypack/models"
	"flypack/service/order"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type OrderHandler interface {
	GetOrders(c *gin.Context)
	PostOrder(c *gin.Context)
	GetOrderByID(c *gin.Context)

}

type orderHandler struct {
	orderService order.OrderService
}

func NewOrderHandler(orderService order.OrderService) OrderHandler{
	return &orderHandler{orderService: orderService}
}

func (h orderHandler) GetOrders(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	filter := c.Query("filter")
	rangeInput := c.Query("range")
	sortInput := c.Query("sort")
	
	log.Printf("Filtros %s", filter)
	log.Printf("Rango %s", rangeInput)
	log.Printf("Ordernar por %s", sortInput)

/* 	body :=&models.GetAllOrderRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	} */
/* 	orders, err :=h.orderService.GetAllOrdersInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, orders) */
		c.IndentedJSON(http.StatusOK, mocks.Orders)
}

func (h orderHandler) PostOrder(c *gin.Context) {
	var body models.RegisterNewOrderRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newOrder, err := h.orderService.CreateOrderInfo(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newOrder )
}

func (h orderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderService.GetOrderInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if order == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"order": order})
}
