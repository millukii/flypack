package handlers

import (
	"flypack/handlers/mocks"
	"flypack/models"
	"flypack/service/shipping"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type ManifiestHandler interface {
	GetManifiests(c *gin.Context)
	PostManifiest(c *gin.Context)
	GetManifiestByID(c *gin.Context)

}

type manifiestHandler struct {
	manifiestService shipping.ManifiestService
}

func NewManifiestHandler(manifiestService shipping.ManifiestService) ManifiestHandler{
	return &manifiestHandler{manifiestService: manifiestService}
}

func (h manifiestHandler) GetManifiests(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	filter := c.Query("filter")
	rangeInput := c.Query("range")
	sortInput := c.Query("sort")
	
	log.Printf("Filtros %s", filter)
	log.Printf("Rango %s", rangeInput)
	log.Printf("Manifiestnar por %s", sortInput)

/* 	body :=&models.GetAllManifiestRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	} */
/* 	manifiests, err :=h.manifiestService.GetAllManifiestsInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, manifiests) */
		c.IndentedJSON(http.StatusOK, mocks.Manifiests)
}

func (h manifiestHandler) PostManifiest(c *gin.Context) {
	var body models.RegisterNewManifiestRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newManifiest, err := h.manifiestService.CreateManifiestInfo(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newManifiest )
}

func (h manifiestHandler) GetManifiestByID(c *gin.Context) {
	id := c.Param("id")

	manifiest, err := h.manifiestService.GetManifiestInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if manifiest == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "manifiest not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"manifiest": manifiest})
}
