package handlers

import (
	"flypack/models"
	"flypack/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type UserStateHandler interface{
	GetUserStates(c *gin.Context)
	PostUserState(c *gin.Context)
	GetUserStateByID(c *gin.Context)
}

type userStatesHandler struct {
	userStateService user.UserStateService
}

func NewUserStateHandler(userStateService user.UserStateService) UserStateHandler{

	return &userStatesHandler{userStateService: userStateService}
}
func (h userStatesHandler) GetUserStates(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	states, err :=h.userStateService.GetUserStates(c)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, states)
}

func (h userStatesHandler) PostUserState(c *gin.Context) {
	var body models.UserState

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	state, err := h.userStateService.AddUserState(c, *&body.State)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, state )
}

func (h userStatesHandler) GetUserStateByID(c *gin.Context) {
	id := c.Param("id")

	state, err := h.userStateService.GetUserState(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if state == ""{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "state not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"state": state})
}
