package migration

import (
	"flypack/config"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)
func buildDBConfig(host, port, user, name, password string, dbType string) *config.DBConfig {
  dbConfig := config.DBConfig{
    Host:     host,
    Port:     port,
    User:     user,
    DBName:   name,
    Password: password,
    DBType:   dbType,
  }
  return &dbConfig
}
func dbURL(dbConfig *config.DBConfig) string {
  if dbConfig.DBType == "cloudsql" {
    return fmt.Sprintf(
      "%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
      dbConfig.User,
      dbConfig.Password,
      dbConfig.Host,
      dbConfig.DBName,
    )
  }
  return fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    dbConfig.User,
    dbConfig.Password,
    dbConfig.Host,
    dbConfig.Port,
    dbConfig.DBName,
  )
}
func main() {
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbPort := os.Getenv("DB_PORT")
  dbHost := os.Getenv("DB_HOST")
  dbName := os.Getenv("DB_NAME")
  dbType := os.Getenv("DB_TYPE")
  fmt.Println("--- Connecting Migrations ---")
  dbConfig := buildDBConfig(dbHost, dbPort, dbUser, dbName, dbPassword, dbType)
  dbURL := dbURL(dbConfig)
  m, err := migrate.New("file://migrations/", dbType+"://"+dbURL)
  
  if err != nil {
    panic(err)
  }
  fmt.Println("--- Running Migration ---")
  m.Steps(10)
}