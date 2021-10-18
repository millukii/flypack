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
	
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (UserRepository, error) {

		return &userRepository{db: db}, nil
}

func(user *userRepository) CreateUser(ctx context.Context,data *models.User) (*models.User, error){
 
query :=	fmt.Sprintf("INSERT INTO %s (user,password,rol_id,user_state_id,people_id) VALUES ('%s','%s',%d,%d,%d)", "users",
			data.User, 	
			data.Password,
			data.Role,
			data.State,
			data.Register,
)
	fmt.Println("Query to execute ", query)
	result, err := user.db.Exec(query)

if err != nil {
	fmt.Printf("could not insert row: %v", err)
	return nil, err
}

// the `Result` type has special methods like `RowsAffected` which returns the
// total number of affected rows reported by the database
// In this case, it will tell us the number of rows that were inserted using
// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return nil, err
	}
// we can log how many rows were inserted
		fmt.Println("inserted", rowsAffected, "rows")

		return data, nil
	}


func(user *userRepository) GetAllUser(ctx context.Context) (*models.AllUserResponse, error){
  rows, err := user.db.Query("SELECT id, user, rol_id,user_state_id,people_id FROM users")
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    // An album slice to hold data from returned rows.
    var users []*models.UserListView

    // Loop through rows, using Scan to assign column data to struct fields.
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
			Count: len(users),
		}, nil
}
func(user *userRepository) GetUser(ctx context.Context, filter, value string) (*models.UserListView, error){
	sqlStatement := fmt.Sprintf("SELECT id, user, rol_id,user_state_id,people_id FROM users WHERE %s=%s;", filter, value)
var res models.UserListView
row := user.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.User, &res.Role,&res.State, &res.Register)
switch err {
case sql.ErrNoRows:
  fmt.Println("No rows were returned!")
  return nil, nil
case nil:
  fmt.Println(user)
default:
  panic(err)
}
return &res, nil
}
func(user *userRepository) UpdateUser(ctx context.Context , data *models.User) (*models.User, error){
 
query :=	fmt.Sprintf("UPDATE  %s set user='%s', password='%s',rol_id=%d,user_state_id=%d,people_id=%d", "users",
			data.User, 	
			data.Password,
			data.Role,
			data.State,
			data.Register,
)
	fmt.Println("Query to execute ", query)
	result, err := user.db.Exec(query)
  
if err != nil {
	fmt.Printf("could not update row: %v", err)
	return nil, err
}

// the `Result` type has special methods like `RowsAffected` which returns the
// total number of affected rows reported by the database
// In this case, it will tell us the number of rows that were inserted using
// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("could not get affected rows: %v", err)
		return nil, err
	}
// we can log how many rows were inserted
		fmt.Println("inserted", rowsAffected, "rows")

		return data, nil
}