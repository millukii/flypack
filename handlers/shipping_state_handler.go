package handlers

import (
	"flypack/models"
	"flypack/service/shipping"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ShippingStateHandler interface{
	GetShippingStates(c *gin.Context)
	PostShippingState(c *gin.Context)
	GetShippingStateByID(c *gin.Context)
}

type shippingStatesHandler struct {
	shippingStateService shipping.ShippingStateService
}

func NewShippingStateHandler(shippingStateService shipping.ShippingStateService) ShippingStateHandler{

	return &shippingStatesHandler{shippingStateService: shippingStateService}
}
func (h shippingStatesHandler) GetShippingStates(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	states, err :=h.shippingStateService.GetShippingStates(c)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, states)
}

func (h shippingStatesHandler) PostShippingState(c *gin.Context) {
	var body models.ShippingState

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	state, err := h.shippingStateService.AddShippingState(c, *&body.State)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, state )
}

func (h shippingStatesHandler) GetShippingStateByID(c *gin.Context) {
	id := c.Param("id")

	state, err := h.shippingStateService.GetShippingState(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if state == ""{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "state not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"state": state})
}
