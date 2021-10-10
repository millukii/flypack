package datasource

import (
	"context"
	"database/sql"
	"fmt"
	"time"

)

type Datasource interface {
	ExecuteQuery(ctx context.Context,query string, destiny interface{}) error
}

type datasource struct {
	db *sql.DB
}

func NewDatasource(ctx context.Context, db *sql.DB) (Datasource, error) {

err := db.PingContext(ctx)
if err != nil {
	return nil, err
}

dbStatus := db.Stats()

fmt.Println("Conexiones en uso ", dbStatus.InUse)
db.SetConnMaxLifetime(time.Minute*5);
db.SetMaxIdleConns(0);
db.SetMaxOpenConns(5);

return &datasource{db: db}, nil
}


func (d *datasource) 	ExecuteQuery(ctx context.Context,query string, destiny interface{}) error {
 rows, err := d.db.QueryContext(ctx, query)

 if err != nil {
	 return err
 }
err = rows.Scan(destiny)
 if err != nil {
	rows.Close()
	 return err
 }
	rows.Close()
	return nil
}