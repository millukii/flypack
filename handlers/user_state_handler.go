package handlers

import (
	"flypack/models"
	"net/http"

	"github.com/gin-gonic/gin"

)


var userStates = []models.UserState{
	{ID: "1", State: "Activo"},
	{ID: "2", State: "InActivo"},
	{ID: "3", State: "Suspendido"},
}

func GetUserStates(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.IndentedJSON(http.StatusOK, userStates)
}

func PostUserStates(c *gin.Context) {
	var newState models.UserState

	if err := c.BindJSON(&newState); err != nil {
		return
	}

	userStates = append(userStates, newState)
	c.IndentedJSON(http.StatusCreated, newState)
}

func GetUserStatesByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range userStates {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user state not found"})
}
