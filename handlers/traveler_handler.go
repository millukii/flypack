package handlers

import (
	"flypack/handlers/mocks"
	"flypack/models"
	"flypack/service/traveler"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type TravelerHandler interface {
	GetTravelers(c *gin.Context)
	PostTraveler(c *gin.Context)
	GetTravelerByID(c *gin.Context)

}

type travelerHandler struct {
	travelerService traveler.TravelerService
}

func NewTravelerHandler(travelerService traveler.TravelerService) TravelerHandler{
	return &travelerHandler{travelerService: travelerService}
}

func (h travelerHandler) GetTravelers(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	filter := c.Query("filter")
	rangeInput := c.Query("range")
	sortInput := c.Query("sort")
	
	log.Printf("Filtros %s", filter)
	log.Printf("Rango %s", rangeInput)
	log.Printf("Travelernar por %s", sortInput)

/* 	body :=&models.GetAllTravelerRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	} */
/* 	travelers, err :=h.travelerService.GetAllTravelersInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, travelers) */
		c.IndentedJSON(http.StatusOK, mocks.Travelers)
}

func (h travelerHandler) PostTraveler(c *gin.Context) {
	var body models.RegisterNewTravelerRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newTraveler, err := h.travelerService.CreateTravelerInfo(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newTraveler )
}

func (h travelerHandler) GetTravelerByID(c *gin.Context) {
	id := c.Param("id")

	traveler, err := h.travelerService.GetTravelerInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if traveler == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "traveler not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"traveler": traveler})
}
