package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"

)

type ShippingStateRepository interface {
	GetShippingState(ctx context.Context, filter, value string) (*models.ShippingState, error)
	GetAllShippingState(ctx context.Context) ([]*models.ShippingState, error)
	CreateShippingState(ctx context.Context, 	user *models.ShippingState) (*models.ShippingState, error)
	UpdateShippingState(
		ctx context.Context,
		user *models.ShippingState)(*models.ShippingState, error)
	DeleteShippingState(ctx context.Context, id string) (error)
}

type shippingStateRepository struct {
	db *sql.DB
	table string
}

func NewShippingStateRepository(db *sql.DB) (ShippingStateRepository, error) {

		return  &shippingStateRepository{db: db, table: "shipping_states"}, nil
}


func(shippingStateRepo *shippingStateRepository) CreateShippingState(ctx context.Context,data *models.ShippingState) (*models.ShippingState, error){

query :=	fmt.Sprintf("INSERT INTO %s (state) VALUES ('%s')", shippingStateRepo.table,
			data.State, 	
)
	fmt.Println("Query to execute ", query)
	result, err := shippingStateRepo.db.Exec(query)

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
func(shippingStateRepo *shippingStateRepository) GetAllShippingState(ctx context.Context) ([]*models.ShippingState, error){


	query :=	fmt.Sprintf("SELECT id, state FROM %s", shippingStateRepo.table)
 rows, err := shippingStateRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var shippingStates []*models.ShippingState

    for rows.Next() {
        var shippingState models.ShippingState
        if err := rows.Scan(&shippingState.ID,
            &shippingState.State); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        shippingStates = append(shippingStates, 
					&models.ShippingState	{
						ID: shippingState.ID, 
						State: shippingState.State,
					})
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return shippingStates, nil
}
func(shippingStateRepo *shippingStateRepository) GetShippingState(ctx context.Context, filter, value string) (*models.ShippingState, error){

	sqlStatement := fmt.Sprintf("SELECT id, state FROM %s WHERE %s=%s;", shippingStateRepo.table,filter, value)
var res models.ShippingState
row := shippingStateRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.State)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}
	return &res, nil
}
func(shippingStateRepo *shippingStateRepository) UpdateShippingState(ctx context.Context , data *models.ShippingState) (*models.ShippingState, error){


query :=	fmt.Sprintf("UPDATE  %s set state='%s'", shippingStateRepo.table,
			data.State, 	
)
	fmt.Println("Query to execute ", query)
	result, err := shippingStateRepo.db.Exec(query)
  
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


func(shippingStateRepo *shippingStateRepository) 	DeleteShippingState(ctx context.Context, id string) (error){
	
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", shippingStateRepo.table,id)

	result, err :=	shippingStateRepo.db.Exec(query)

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