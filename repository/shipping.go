package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"

)

type ShippingRepository interface {
	GetShipping(ctx context.Context, filter, value string) (*models.Shipping, error)
	GetAllShipping(ctx context.Context) ([]*models.Shipping, error)
	CreateShipping(ctx context.Context, 	shippingRepo *models.Shipping) (*models.Shipping, error)
	UpdateShipping(
		ctx context.Context,
		shippingRepo *models.Shipping)(*models.Shipping, error)
	DeleteShipping(ctx context.Context, id string) (error)
}

type shippingRepository struct {
	db *sql.DB
	table string
}

func NewShippingRepository(	db *sql.DB) (ShippingRepository, error) {

		return  &shippingRepository{db: db, table: "shipping"}, nil
}


func(shippingRepo *shippingRepository) CreateShipping(ctx context.Context,data *models.Shipping) (*models.Shipping, error){

query :=	fmt.Sprintf("INSERT INTO %s (order_nro,boleta_nro,type,value,date,shipping_states_id,companies_id,delivery_id,client_id) VALUES (%d','%s',%s,%f,%s,%d,%d,%d,%d)", shippingRepo.table,
			data.OrderNumber, 	
			data.TickerNumber,
			data.ShippingType,
			data.Value,
			data.Date,
			data.ShippingState, 
			data.Company,
			data.Delivery,
			data.Client,
)
	fmt.Println("Query to execute ", query)
	result, err := shippingRepo.db.Exec(query)

if err != nil {
	fmt.Printf("could not insert row: %v", err)
	return nil, err
}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return nil, err
	}
	fmt.Println("inserted", rowsAffected, "rows")

		return data, nil
}
func(shippingRepo *shippingRepository) GetAllShipping(ctx context.Context) ([]*models.Shipping, error){

	query := fmt.Sprintf("SELECT id,order_nro,boleta_nro,type,value,date,shipping_states_id,companies_id,delivery_id,client_id from %s", shippingRepo.table)
rows, err := shippingRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var shippings []*models.Shipping
    for rows.Next() {
        var shipping models.Shipping
        if err := rows.Scan(&shipping.ID, &shipping.OrderNumber,&shipping.TickerNumber,
            &shipping.ShippingType, &shipping.Value,&shipping.Date,&shipping.ShippingState,&shipping.Company, &shipping.Delivery, &shipping.Client); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        shippings = append(shippings, &shipping)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return shippings, nil
}
func(shippingRepo *shippingRepository) GetShipping(ctx context.Context, filter, value string) (*models.Shipping, error){

	sqlStatement := fmt.Sprintf("SELECT id,order_nro,boleta_nro,type,value,date,shipping_states_id,companies_id,delivery_id,client_id FROM %s WHERE %s=%s;", 
	shippingRepo.table,filter, value)
var res models.Shipping
row := shippingRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.OrderNumber, &res.TickerNumber,&res.ShippingType, &res.Value,&res.Date, &res.ShippingState, &res.Company, &res.Delivery, &res.Client)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}
	return &res, nil
}
func(shippingRepo *shippingRepository) UpdateShipping(ctx context.Context , data *models.Shipping) (*models.Shipping, error){

query :=	fmt.Sprintf("UPDATE  %s set order_nro='%d', boleta_nro='%s',type=%s,value=%f,date=%s, shipping_states_id=%d,companies_id=%d,delivery_id=%d,client_id=%d", shippingRepo.table,
			data.OrderNumber, 	
			data.TickerNumber,
			data.ShippingType,
			data.Value,
			data.Date,
			data.ShippingState,
			data.Company, 
			data.Delivery,
			data.Client,
)
	fmt.Println("Query to execute ", query)
	result, err := shippingRepo.db.Exec(query)
  
if err != nil {
	fmt.Printf("could not update row: %v", err)
	return nil, err
}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return nil, err
	}
	fmt.Println("inserted", rowsAffected, "rows")

		return data, nil
}

func(shippingRepo *shippingRepository) DeleteShipping(ctx context.Context, id string) (error) {

		
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", shippingRepo.table,id)

	result, err :=	shippingRepo.db.Exec(query)

if err != nil {
	fmt.Printf("could not deleted row: %v", err)
	return err
}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return err
	}
 fmt.Println("deleted", rowsAffected, "rows")

	return nil
}