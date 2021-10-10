package handlers

import (
	"flypack/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


var shippingStates = []models.ShippingState{
	{ID: "1", State: "BlueTrain"},
}

func GetShippingStates(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.IndentedJSON(http.StatusOK, shippingStates)
}

func PostShippingState(c *gin.Context) {
	var newShippingState models.ShippingState

	if err := c.BindJSON(&newShippingState); err != nil {
		return
	}

	shippingStates = append(shippingStates, newShippingState)
	c.IndentedJSON(http.StatusCreated, newShippingState)
}

func GetShippingStateByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range shippingStates {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ShippingState not found"})
}
