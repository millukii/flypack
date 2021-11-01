package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"

)

type RoleRepository interface {
	GetRole(ctx context.Context, filter, value string) (*models.Role, error)
	GetAllRole(ctx context.Context) ([]*models.Role, error)
	CreateRole(ctx context.Context, 	user *models.Role) (*models.Role, error)
	UpdateRole(
		ctx context.Context,
		user *models.Role)(*models.Role, error)
	DeleteRole(ctx context.Context, id string) (error)
}

type roleRepository struct {
	db *sql.DB
	table string
}

func NewRoleRepository(db *sql.DB) (RoleRepository, error) {

		return  &roleRepository{db: db, table: "roles"}, nil
}


func(roleRepo *roleRepository) CreateRole(ctx context.Context,data *models.Role) (*models.Role, error){


query :=	fmt.Sprintf("INSERT INTO %s (state) VALUES ('%s')", roleRepo.table,
			data.Role, 	
)
	fmt.Println("Query to execute ", query)
	result, err := roleRepo.db.Exec(query)

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
func(roleRepo *roleRepository) GetAllRole(ctx context.Context) ([]*models.Role, error){


	query :=	fmt.Sprintf("SELECT id, role FROM %s", roleRepo.table)
 rows, err := roleRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var roles []*models.Role

    for rows.Next() {
        var role models.Role
        if err := rows.Scan(&role.ID,
            &role.Role); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        roles = append(roles, 
					&models.Role	{
						ID: role.ID, 
						Role: role.Role,
					})
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return roles, nil
}
func(roleRepo *roleRepository) GetRole(ctx context.Context, filter, value string) (*models.Role, error){

	sqlStatement := fmt.Sprintf("SELECT id, role FROM %s WHERE %s=%s;", roleRepo.table,filter, value)
var res models.Role
row := roleRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.Role)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}
	return &res, nil
}
func(roleRepo *roleRepository) UpdateRole(ctx context.Context , data *models.Role) (*models.Role, error){



query :=	fmt.Sprintf("UPDATE  %s set role='%s'", roleRepo.table,
			data.Role, 	
)
	fmt.Println("Query to execute ", query)
	result, err := roleRepo.db.Exec(query)
  
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


func(roleRepo *roleRepository)	DeleteRole(ctx context.Context, id string) (error){

	
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", roleRepo.table,id)

	result, err :=	roleRepo.db.Exec(query)

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
