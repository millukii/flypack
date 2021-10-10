package mysql

import (
	"database/sql"
)

func MySQLConnect(dns string) (*sql.DB,error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil,err
	}
	return db, nil
}