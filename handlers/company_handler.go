package handlers

import (
	"flypack/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

)


var companies = []models.Company{
	{ID: "1", Rut: 22221960, DV: "4", Razon: "Ventas", Address: "Los alerces 2363", City: "Santiago", Commune: "Nunoa", LegalPerson: 1},
		{ID: "2", Rut: 3423432432, DV: "4", Razon: "Empresa3", Address: "cualqioer address", City: "Santiago", Commune: "Nunoa", LegalPerson: 3},
}

func GetCompanies(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	filter := c.Query("filter")
	rangeInput := c.Query("range")
	sortInput := c.Query("sort")
	
	log.Printf("Filtros %s", filter)
	log.Printf("Rango %s", rangeInput)
	log.Printf("Ordernar por %s", sortInput)

	c.IndentedJSON(http.StatusOK, companies)
}

func PostCompany(c *gin.Context) {
	var newCompany models.Company

	if err := c.BindJSON(&newCompany); err != nil {
		return
	}

	companies = append(companies, newCompany)
	c.IndentedJSON(http.StatusCreated, newCompany)
}

func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range companies {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "company not found"})
}
