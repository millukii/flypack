package handlers

import (
	"flypack/models"
	"flypack/service/user"
	"net/http"

	"github.com/gin-gonic/gin"

)
type RoleHandler interface{
	GetRoles(c *gin.Context)
	PostRol(c *gin.Context)
	 GetRolByID(c *gin.Context)
}

type roleHandler struct {
	roleService user.RoleService
}

func NewRoleHandler(	roleService user.RoleService) RoleHandler {
	return &roleHandler{
		roleService: roleService,
	}
}

func (r roleHandler) GetRoles(c *gin.Context) {
	//Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "Content-Type")

	roles, err :=	r.roleService.GetRoles(c)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, roles)
}

func (r roleHandler) PostRol(c *gin.Context) {
	var newRole models.Role

	if err := c.BindJSON(&newRole); err != nil {
		return
	}
	role, err :=	r.roleService.AddRole(c, newRole.Role)
	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}
	c.IndentedJSON(http.StatusOK, role)
}

func (r roleHandler) GetRolByID(c *gin.Context) {
	id := c.Param("id")

	role, err := r.roleService.GetRole(c, id)

	if err != nil {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error in service"})
	}

	if role == ""{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "role not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"role": role})
}
