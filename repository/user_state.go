package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"

)

type UserStateRepository interface {
	GetUserState(ctx context.Context, filter, value string) (*models.UserState, error)
	GetAllUserState(ctx context.Context) ([]*models.UserState, error)
	CreateUserState(ctx context.Context, 	user *models.UserState) (*models.UserState, error)
	UpdateUserState(
		ctx context.Context,
		user *models.UserState)(*models.UserState, error)
	DeleteUserState(ctx context.Context, id string) (error)
}

type userStateRepository struct {
 db *sql.DB
 table string
}

func NewUserStateRepository( db *sql.DB) (UserStateRepository, error) {

		return  &userStateRepository{db: db, table: "users_state"}, nil
}


func(userStateRepo *userStateRepository) CreateUserState(ctx context.Context,data *models.UserState) (*models.UserState, error){

query :=	fmt.Sprintf("INSERT INTO %s (state) VALUES ('%s')", userStateRepo.table,
			data.State, 	
)
	fmt.Println("Query to execute ", query)
	result, err := userStateRepo.db.Exec(query)

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
func(userStateRepo *userStateRepository) GetAllUserState(ctx context.Context) ([]*models.UserState, error){

	query :=	fmt.Sprintf("SELECT id, state FROM %s", userStateRepo.table)
 rows, err := userStateRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var users []*models.UserState

    for rows.Next() {
        var user models.UserState
        if err := rows.Scan(&user.ID,
            &user.State); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        users = append(users, 
					&models.UserState	{
						ID: user.ID, 
						State: user.State,
					})
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return users, nil
}
func(userStateRepo *userStateRepository) GetUserState(ctx context.Context, filter, value string) (*models.UserState, error){

	sqlStatement := fmt.Sprintf("SELECT id, state FROM %s WHERE %s=%s;", userStateRepo.table,filter, value)
var res models.UserState
row := userStateRepo.db.QueryRow(sqlStatement)
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
func(userStateRepo *userStateRepository) UpdateUserState(ctx context.Context , data *models.UserState) (*models.UserState, error){

query :=	fmt.Sprintf("UPDATE  %s set state='%s'", userStateRepo.table,
			data.State, 	
)
	fmt.Println("Query to execute ", query)
	result, err := userStateRepo.db.Exec(query)
  
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


func(userStateRepo *userStateRepository) 	DeleteUserState(ctx context.Context, id string) (error){
	
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", userStateRepo.table,id)

	result, err :=	userStateRepo.db.Exec(query)

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