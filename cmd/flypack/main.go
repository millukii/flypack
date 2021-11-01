package main

import (
	"flypack/config"
	"flypack/db/mysqldts"
	"flypack/handlers"
	"flypack/logger"
	"flypack/repository"
	"flypack/service/company"
	"flypack/service/people"
	"flypack/service/shipping"
	"flypack/service/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"

)

func init() {
	    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("unable to load .env file")
    }
}
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
db,err := mysqldts.MySQLConnect(dns)

if err != nil {
	fmt.Println("Error connection database: ", err.Error())
		panic(err)
}

if db ==nil {
		fmt.Println("Error connection database")
		panic("Driver error")
}
connectionError := db.Ping()
if connectionError  != nil {
		fmt.Println("Error ping database", connectionError.Error())
		panic(connectionError.Error())
}

logger,_ := logger.NewApplicationLogger()
userRepo, err := repository.NewUserRepository(db)
if err != nil {
		fmt.Println("Error creating user repository: ", err.Error())
		panic(err)
}

userService := user.NewUserService(userRepo, logger)
accountService := user.NewUserAccountService(userRepo)
passwordGenerator := user.PasswordGenerator{}
userHandler := handlers.NewUserHandler(userService, accountService,passwordGenerator)

router.GET("/users", userHandler.GetUsers)
router.GET("/users/:id", userHandler.GetUserByID)
router.POST("/users", userHandler.PostUsers)
router.POST("/users/login", userHandler.Login)


userStateRepo, err := repository.NewUserStateRepository(db)
userStateService := user.NewUserStateService(userStateRepo)
userStatesHandler := handlers.NewUserStateHandler(userStateService)

router.GET("/userstates", userStatesHandler.GetUserStates)
router.GET("/userstates/:id", userStatesHandler.GetUserStateByID)
router.POST("/userstates", userStatesHandler.PostUserState)

companyRepo, err := repository.NewCompanyRepository(db)
companyService := company.NewCompanyService(companyRepo)
companyHandler := handlers.NewCompanyHandler(companyService)

router.GET("/companies",companyHandler.GetCompanies)
router.GET("/companies/:id", companyHandler.GetCompanyByID)
router.POST("/companies", companyHandler.PostCompany)

peopleRepo, err := repository.NewPeopleRepository(db)
peopleService := people.NewPeopleService(peopleRepo)
peopleHandler := handlers.NewPeopleHandler(peopleService)

router.GET("/people", peopleHandler.GetPeople)
router.GET("/people/:id", peopleHandler.GetPeopleByID)
router.POST("/people", peopleHandler.PostPeople)


	roleRepository, err := repository.NewRoleRepository(db)
	roleService := user.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	router.GET("/roles", roleHandler.GetRoles)
	router.GET("/roles/:id", roleHandler.GetRolByID)
	router.POST("/roles", roleHandler.PostRol)

	
	shippingRepo, err := repository.NewShippingRepository(db)
	shippingService := shipping.NewShippingService(shippingRepo)
	shippingsHandler := handlers.NewShippingHandler(shippingService)
	router.GET("/shippings", shippingsHandler.GetShippings)
	router.GET("/shippings/:id", shippingsHandler.GetShippings)
	router.POST("/shippings", shippingsHandler.PostShippings)


shippingStatesRepo, err := repository.NewShippingStateRepository(db)
shippingStatesService := shipping.NewShippingStateService(shippingStatesRepo)
shippingStatesHandler := handlers.NewShippingStateHandler(shippingStatesService)

	router.GET("/shippingstates", shippingStatesHandler.GetShippingStates)
	router.GET("/shippingstates/:id", shippingStatesHandler.GetShippingStateByID)
	router.POST("/shippingstates", shippingStatesHandler.PostShippingState)

	
    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()

}
