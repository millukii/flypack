package config

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{})
}

func Connect(connection string) (*sql.DB,error) {

	return	 sql.Open("mysql", connection)

}