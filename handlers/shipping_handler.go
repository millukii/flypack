package handlers

import (
	"flypack/models"
	"flypack/service/shipping"
	"net/http"

	"github.com/gin-gonic/gin"

)

type ShippingHandler interface {
	GetShippings(c *gin.Context)
	PostShippings(c *gin.Context)
	GetShippingByID(c *gin.Context)
}
type shippingHandler struct {
	shippingService shipping.ShippingService
}

func NewShippingHandler(shippingService shipping.ShippingService) (ShippingHandler, error) {

	return &shippingHandler{
			shippingService: shippingService,
	}, nil
}
var shippings = []models.Shipping{
	{ID: "1", OrderNumber: 1, TickerNumber: "1", ShippingType: "1", Value: 1, Date: "2015-01-01", ShippingState: 1, Company: 1, Delivery: 1, Client: 1},
}

func (h shippingHandler) GetShippings(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.IndentedJSON(http.StatusOK, shippings)
}

func (h shippingHandler) PostShippings(c *gin.Context) {
	var newShipping models.Shipping

	if err := c.BindJSON(&newShipping); err != nil {
		return
	}

	shippings = append(shippings, newShipping)
	c.IndentedJSON(http.StatusCreated, newShipping)
}

func (h shippingHandler) GetShippingByID(c *gin.Context) {
	id := c.Param("id")

	shipping, err := h.shippingService.GetShippingById(c, id)


	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if shipping == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "shipping not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"shipping": shipping})
}
