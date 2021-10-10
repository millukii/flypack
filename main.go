package main

import (
	"flypack/handlers"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()

router.Use(cors.Middleware(cors.Config{
	Origins:        "*",
	Methods:        "GET, PUT, POST, DELETE, PATCH, OPTIONS",
	RequestHeaders: "Origin, Authorization, Content-Type",
	ExposedHeaders: "",
	MaxAge: 50 * time.Second,
	Credentials: true,
	ValidateHeaders: false,
}))

router.Use(gin.Logger())


	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.POST("/users", handlers.PostUsers)
		router.POST("/users/login", handlers.Login)

	router.GET("/userstates", handlers.GetUserStates)
	router.GET("/userstates/:id", handlers.GetUserStatesByID)
	router.POST("/userstates", handlers.PostUserStates)

	router.GET("/companies",handlers.GetCompanies)
	router.GET("/companies/:id", handlers.GetCompanyByID)
	router.POST("/companies", handlers.PostCompany)

	router.GET("/people", handlers.GetPeople)
	router.GET("/people/:id", handlers.GetPeopleByID)
	router.POST("/people", handlers.PostPeople)

	router.GET("/roles", handlers.GetRoles)
	router.GET("/roles/:id", handlers.GetRolByID)
	router.POST("/roles", handlers.PostRol)

	router.GET("/shippings", handlers.GetShippings)
	router.GET("/shippings/:id", handlers.GetShippings)
	router.POST("/shippings", handlers.PostShippings)

	router.GET("/shippingstates", handlers.GetShippingStates)
	router.GET("/shippingstates/:id", handlers.GetShippingStateByID)
	router.POST("/shippingstates", handlers.PostShippingState)

	router.Run("localhost:8080")
}
