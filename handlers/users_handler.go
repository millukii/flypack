package handlers

import (
	"flypack/models"
	"flypack/service/user"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

)

type UserHandler interface {
	GetUsers(c *gin.Context)
	Login(c *gin.Context)
	PostUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
}

type userHandler struct {
	userService user.UserService
	accountService user.UserAccountService

}
func NewUserHandler(userService user.UserService, 	accountService user.UserAccountService) UserHandler {
	return &userHandler{userService: userService,	accountService: accountService}
}

func (h userHandler) GetUsers(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")
	body :=&models.GetAllUserRequest{}
	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
	users, err :=h.userService.GetAllUsersInfo(c, body)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (h userHandler) Login(c *gin.Context) {
  var user models.User
  if err := c.ShouldBindJSON(&user); err != nil {
     c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
     return
  }
  //compare the user from the request, with the one we defined:
  if user.User !="admin" || user.Password != "123" {
     c.JSON(http.StatusUnauthorized, "Please provide valid login details")
     return
  }

	id, _ := strconv.ParseUint(user.ID, 0, 64) 
  token, err := CreateToken(id)

  if err != nil {
     c.JSON(http.StatusUnprocessableEntity, err.Error())
     return
  }
  c.JSON(http.StatusOK, token)
}

func (h userHandler) PostUsers(c *gin.Context) {
	var body models.RegisterNewUserRequest

	if err := c.ShouldBindBodyWith(&body,binding.JSON);err!=nil{
		c.AbortWithError(http.StatusBadRequest,err)
		return
	}
 
	newUserResponse, err := h.accountService.RegisterNewAccount(c, &body)
	
		if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	c.IndentedJSON(http.StatusCreated, newUserResponse )
}

func (h userHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetUserInfo(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if user == nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

func CreateToken(userid uint64) (string, error) {
  var err error
  //Creating Access Token
  os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }
  return token, nil
}