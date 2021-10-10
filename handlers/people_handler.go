package handlers

import (
	"flypack/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



var peopleList = []models.People{
	{ID: "1", Rut: 22221960, DV: "4", Name: "Ventas", Lastname: "Ventas", Address: "Los alerces 2363", City: "Santiago", Commune: "Nunoa", Email: "jojo@mail.com", Phone: "3435435", Profile: 1},
		{ID: "2", Rut: 2343453, DV: "4", Name: "Vendedor2", Lastname: "apellido3", Address: "monjitas 744", City: "Santiago", Commune: "Santiago Centro", Email: "jojo@mail.com", Phone: "3435435", Profile: 2},
}

func GetPeople(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.IndentedJSON(http.StatusOK, peopleList)
}

func PostPeople(c *gin.Context) {
	var newPeople models.People

	if err := c.BindJSON(&newPeople); err != nil {
		return
	}

	peopleList = append(peopleList, newPeople)
	c.IndentedJSON(http.StatusCreated, newPeople)
}

func GetPeopleByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range peopleList {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "People not found"})
}
