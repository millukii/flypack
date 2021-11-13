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

type ShipmentPackageHandler interface {
	GetShipmentPackages(c *gin.Context)
	PostShipmentPackage(c *gin.Context)
	GetShipmentPackageByID(c *gin.Context)

}

type shipmentPackageHandler struct {
	shipmentPackageService shipping.ShipmentPackagesService
}

func NewShipmentPackageHandler(shipmentPackageService shipping.ShipmentPackagesService) ShipmentPackageHandler{
	return &shipmentPackageHandler{shipmentPackageService: shipmentPackageService}
}

func (h shipmentPackageHandler) GetShipmentPackages(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	filter := c.Query("filter")
	rangeInput := c.Query("range")
	sortInput := c.Query("sort")
	
	log.Printf("Filtros %s", filter)
	log.Printf("Rango %s", rangeInput)
	log.Printf("ShipmentPackagenar por %s", sortInput)

/* 	body :=&models.GetAllShipmentPackageRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	} */
/* 	shipmentPackages, err :=h.shipmentPackageService.GetAllShipmentPackagesInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, shipmentPackages) */
		c.IndentedJSON(http.StatusOK, mocks.ShipmentPackages)
}

func (h shipmentPackageHandler) PostShipmentPackage(c *gin.Context) {
	var body models.RegisterNewShipmentPackagesRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newShipmentPackage, err := h.shipmentPackageService.CreateShipmentPackagesInfo(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newShipmentPackage )
}

func (h shipmentPackageHandler) GetShipmentPackageByID(c *gin.Context) {
	id := c.Param("id")

	shipmentPackage, err := h.shipmentPackageService.GetShipmentPackagesInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if shipmentPackage == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "shipmentPackage not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"shipmentPackage": shipmentPackage})
}
