package mysqldts

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnect(dns string) (*sql.DB,error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil,err
	}
	return db, nil
}