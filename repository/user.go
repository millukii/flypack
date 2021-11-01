package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"

)

type UserRepository interface {
	GetUser(ctx context.Context, filter, value string) (*models.UserListView, error)
	GetAllUser(ctx context.Context) (*models.AllUserResponse, error)
	CreateUser(ctx context.Context, 	user *models.User) (*models.User, error)
	UpdateUser(
		ctx context.Context,
		user *models.User)(*models.User, error)
	DeleteUser(ctx context.Context, id string) (error)
}

type userRepository struct {
	db *sql.DB
	table string
}

func NewUserRepository(db *sql.DB) (UserRepository, error) {

		return &userRepository{db: db, table: "user"}, nil
}

func(userRepo *userRepository) CreateUser(ctx context.Context,data *models.User) (*models.User, error){
 
query :=	fmt.Sprintf("INSERT INTO %s (user,password,rol_id,user_state_id,people_id) VALUES ('%s','%s',%d,%d,%d)", userRepo.table,
			data.User, 	
			data.Password,
			data.Role,
			data.State,
			data.Register,
)
	fmt.Println("Query to execute ", query)
	result, err := userRepo.db.Exec(query)

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


func(userRepo *userRepository) GetAllUser(ctx context.Context) (*models.AllUserResponse, error){

	query := fmt.Sprintf("SELECT id, user, rol_id,user_state_id,people_id FROM %s", userRepo.table)
  rows, err := userRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var users []*models.UserListView
    for rows.Next() {
        var user models.UserListView
        if err := rows.Scan(&user.ID, &user.User,&user.Role,
            &user.State, &user.Register); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        users = append(users, 
					&models.UserListView	{
						ID: user.ID, 
						User: user.User,
						Role: user.Role,
						State: user.State,
						Register: user.Register,
					})
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return &models.AllUserResponse{
			UserListView: users,
			Total: len(users),
		}, nil
}
func(userRepo *userRepository) GetUser(ctx context.Context, filter, value string) (*models.UserListView, error){
	sqlStatement := fmt.Sprintf("SELECT id, user, rol_id,user_state_id,people_id FROM %s WHERE %s=%s;", userRepo.table,filter, value)
var res models.UserListView
row := userRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.User, &res.Role,&res.State, &res.Register)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}
	return &res, nil
}
func(userRepo *userRepository) UpdateUser(ctx context.Context , data *models.User) (*models.User, error){
 
query :=	fmt.Sprintf("UPDATE  %s set user='%s', password='%s',rol_id=%d,user_state_id=%d,people_id=%d", userRepo.table,
			data.User, 	
			data.Password,
			data.Role,
			data.State,
			data.Register,
)
	fmt.Println("Query to execute ", query)
	result, err := userRepo.db.Exec(query)
  
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

func(userRepo *userRepository) 	DeleteUser(ctx context.Context, id string) (error){
	
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", userRepo.table,id)

	result, err :=	userRepo.db.Exec(query)

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