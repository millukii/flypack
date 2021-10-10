package handlers

import (
	"flypack/models"
	"flypack/service/shipping"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type ShippingHandler interface {
	GetShippings(c *gin.Context)
	PostShippings(c *gin.Context)
	GetShippingByID(c *gin.Context)
}
type shippingHandler struct {
	shippingService shipping.ShippingService
}

func NewShippingHandler(shippingService shipping.ShippingService) (ShippingHandler) {

	return &shippingHandler{
			shippingService: shippingService,
	}
}

func (h shippingHandler) GetShippings(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	body :=&models.GetAllShippingRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	users, err :=h.shippingService.GetAllShippingInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (h shippingHandler) PostShippings(c *gin.Context) {
	var body models.CreateShippingRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newShipping, err := h.shippingService.CreateShipping(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newShipping )
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
