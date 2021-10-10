package main

import (
	"flypack/config"
	"flypack/db/mysql"
	"flypack/handlers"
	"flypack/repository"
	"flypack/service/company"
	"flypack/service/people"
	"flypack/service/shipping"
	"flypack/service/user"
	"fmt"
	"os"
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


dbconfig := &config.DBConfig{
	Host: os.Getenv("DB_HOST"),
	Port: os.Getenv("DB_PORT"),
	DBName: os.Getenv("DB_NAME"),
	User: os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	
}

dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
dbconfig.User, dbconfig.Password,dbconfig.Host, dbconfig.Port, dbconfig.DBName)
db,err := mysql.MySQLConnect(dns)

if err != nil {

}
userRepo, err := repository.NewUserRepository(db)
if err != nil {
	
}
userService, err := user.NewUserService(&userRepo)
accountService, err := user.NewUserAccountService(&userRepo)
userHandler := handlers.NewUserHandler(userService, accountService)

router.GET("/users", userHandler.GetUsers)
router.GET("/users/:id", userHandler.GetUserByID)
router.POST("/users", userHandler.PostUsers)
router.POST("/users/login", userHandler.Login)


userStateRepo, err := repository.NewUserStateRepository(db)
userStateService, err := user.NewUserStateService(userStateRepo)
userStatesHandler := handlers.NewUserStateHandler(userStateService)

router.GET("/userstates", userStatesHandler.GetUserStates)
router.GET("/userstates/:id", userStatesHandler.GetUserStateByID)
router.POST("/userstates", userStatesHandler.PostUserState)

companyRepo, err := repository.NewCompanyRepository(db)
companyService, err := company.NewCompanyService(companyRepo)
companyHandler := handlers.NewCompanyHandler(companyService)

router.GET("/companies",companyHandler.GetCompanies)
router.GET("/companies/:id", companyHandler.GetCompanyByID)
router.POST("/companies", companyHandler.PostCompany)

peopleRepo, err := repository.NewPeopleRepository(db)
peopleService, err := people.NewPeopleService(peopleRepo)
peopleHandler := handlers.NewPeopleHandler(peopleService)

router.GET("/people", peopleHandler.GetPeople)
router.GET("/people/:id", peopleHandler.GetPeopleByID)
router.POST("/people", peopleHandler.PostPeople)


	roleRepository, err := repository.NewRoleRepository(db)
	roleService, err := user.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	router.GET("/roles", roleHandler.GetRoles)
	router.GET("/roles/:id", roleHandler.GetRolByID)
	router.POST("/roles", roleHandler.PostRol)

	
	shippingRepo, err := repository.NewShippingRepository(db)
	shippingService, err := shipping.NewShippingService(&shippingRepo)
	shippingsHandler := handlers.NewShippingHandler(shippingService)
	router.GET("/shippings", shippingsHandler.GetShippings)
	router.GET("/shippings/:id", shippingsHandler.GetShippings)
	router.POST("/shippings", shippingsHandler.PostShippings)


shippingStatesRepo, err := repository.NewShippingStateRepository(db)
shippingStatesService, err := shipping.NewShippingStateService(shippingStatesRepo)
shippingStatesHandler := handlers.NewShippingStateHandler(shippingStatesService)

	router.GET("/shippingstates", shippingStatesHandler.GetShippingStates)
	router.GET("/shippingstates/:id", shippingStatesHandler.GetShippingStateByID)
	router.POST("/shippingstates", shippingStatesHandler.PostShippingState)

	router.Run("localhost:8080")
}
