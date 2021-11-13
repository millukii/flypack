package repository

import (
	"context"
	"database/sql"
	"flypack/models"

)

type OrderRepository interface {
	GetOrder(ctx context.Context, filter, value string) (*models.Order, error)
	GetAllOrder(ctx context.Context) ([]*models.Order, error)
	CreateOrder(ctx context.Context, 	user *models.Order) (*models.Order, error)
	UpdateOrder(
		ctx context.Context,
		user *models.Order)(*models.Order, error)
	DeleteOrder(ctx context.Context, id string) (error)
}

type orderRepository struct {
	db *sql.DB
	table string
}

func NewOrderRepository(db *sql.DB) (OrderRepository, error) {

		return &orderRepository{db: db, table: "companies"}, nil
}

func(orderRepo *orderRepository) CreateOrder(ctx context.Context,data *models.Order) (*models.Order, error){
/* 
query :=	fmt.Sprintf("INSERT INTO %s (rut, dv, razon, address, city, commune, people_id) VALUES (%d','%s',%s,%s,%s,%s,%d )", orderRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,
)
	fmt.Println("Query to execute ", query)
	result, err := orderRepo.db.Exec(query)

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
 */
		return nil, nil
}
func(orderRepo *orderRepository) GetAllOrder(ctx context.Context) ([]*models.Order, error){

/* 
	query := fmt.Sprintf("SELECT  id, rut, dv, razon, address, city, commune, people_id from %s", orderRepo.table)
rows, err := orderRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var orderRecords []*models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.ID, &order.Rut,&order.DV,
            &order.Razon, &order.Address,&order.City, &order.Commune, &order.LegalPerson); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        orderRecords = append(orderRecords, &order)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    } */
    return nil, nil
}
func(orderRepo *orderRepository) GetOrder(ctx context.Context, filter, value string) (*models.Order, error){
/* 
	sqlStatement := fmt.Sprintf("SELECT id, rut, dv, razon, address, city, commune, people_id FROM %s WHERE %s=%s;", orderRepo.table,filter, value)
var res models.Order
row := orderRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.Rut, &res.DV,&res.Razon, &res.Address,&res.City, &res.Commune,&res.LegalPerson)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}*/
	return nil, nil
}

func(orderRepo *orderRepository) UpdateOrder(ctx context.Context , data *models.Order) (*models.Order, error){
/*
query :=	fmt.Sprintf("UPDATE  %s set rut='%d', dv='%s',razon=%s, address=%s,city=%s,commune=%s,people_id=%d",orderRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,

)
	fmt.Println("Query to execute ", query)
	result, err := orderRepo.db.Exec(query)
  
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
 */
		return nil, nil
}
func(orderRepo *orderRepository) 	DeleteOrder(ctx context.Context, id string) (error){
	/* 
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", orderRepo.table,id)

	result, err :=	orderRepo.db.Exec(query)

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
 */
	return nil
}