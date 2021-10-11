package handlers

import (
	"flypack/models"
	"flypack/service/people"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)


type PeopleHandler interface{
	GetPeople(c *gin.Context)
	PostPeople(c *gin.Context) 
	GetPeopleByID(c *gin.Context)
}

type peopleHandler struct {
	peopleService people.PeopleService
}

func NewPeopleHandler(peopleService people.PeopleService) PeopleHandler{
	
	return &peopleHandler{
		peopleService: peopleService,
	}
}

func (h peopleHandler) GetPeople(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	body :=&models.GetAllPeopleRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	people, err :=h.peopleService.GetAllPeoplesInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, people)
}

func (h peopleHandler) PostPeople(c *gin.Context) {
	var body models.RegisterNewPeopleRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newRegister, err := h.peopleService.RegisterPerson(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newRegister )
}

func (h peopleHandler) GetPeopleByID(c *gin.Context) {
	id := c.Param("id")

	register, err := h.peopleService.GetPeopleInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if register == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "people register not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"people": register})
}
