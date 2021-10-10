package handlers

import (
	"flypack/models"
	"flypack/service/company"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type CompanyHandler interface {
	GetCompanies(c *gin.Context)
	PostCompany(c *gin.Context)
	GetCompanyByID(c *gin.Context)

}

type companyHandler struct {
	companyService company.CompanyService
}

func NewCompanyHandler(companyService company.CompanyService) CompanyHandler{
	return &companyHandler{companyService: companyService}
}

func (h companyHandler) GetCompanies(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	filter := c.Query("filter")
	rangeInput := c.Query("range")
	sortInput := c.Query("sort")
	
	log.Printf("Filtros %s", filter)
	log.Printf("Rango %s", rangeInput)
	log.Printf("Ordernar por %s", sortInput)

	body :=&models.GetAllCompanyRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	companies, err :=h.companyService.GetAllCompanysInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, companies)
}

func (h companyHandler) PostCompany(c *gin.Context) {
	var body models.RegisterNewCompanyRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newCompany, err := h.companyService.CreateCompanyInfo(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newCompany )
}

func (h companyHandler) GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	company, err := h.companyService.GetCompanyInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if company == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "company not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"company": company})
}
