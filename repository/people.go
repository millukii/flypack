package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"

)

type PeopleRepository interface {
	GetPeople(ctx context.Context, filter, value string) (*models.People, error)
	GetAllPeople(ctx context.Context) ([]*models.People, error)
	CreatePeople(ctx context.Context, 	people *models.People) (*models.People, error)
	UpdatePeople(
		ctx context.Context,
		people *models.People)(*models.People, error)
	DeletePeople(ctx context.Context, id string) (error)
}

type peopleRepository struct {
	db *sql.DB
	table string
}

func NewPeopleRepository(	db *sql.DB) (PeopleRepository, error) {

		return  &peopleRepository{db: db, table: "people",}, nil
}

func(peopleRepo *peopleRepository) CreatePeople(ctx context.Context,data *models.People) (*models.People, error){


query :=	fmt.Sprintf("INSERT INTO %s (rut,dv,name,lastname,address,city,commune,email,phone,profile_id) VALUES (%d','%s',%s,%s,%s,%s,%s,%s,%s,%d)", peopleRepo.table,
			data.Rut, 	
			data.DV,
			data.Name,
			data.Lastname,
			data.Address,
			data.City,
			data.Commune,
			data.Email,
			data.Phone,
			data.Profile,
)
	fmt.Println("Query to execute ", query)
	result, err := peopleRepo.db.Exec(query)

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
func(peopleRepo *peopleRepository) GetAllPeople(ctx context.Context) ([]*models.People, error){


	query := fmt.Sprintf("SELECT  id ,rut,dv,name,lastname,address,city,commune,email,phone,profile_id from %s", peopleRepo.table)
rows, err := peopleRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var peopleRecords []*models.People
    for rows.Next() {
        var people models.People
        if err := rows.Scan(&people.ID, &people.Rut,&people.DV,
            &people.Name, &people.Lastname,&people.Address,&people.City, &people.Commune, &people.Email, &people.Phone, &people.Profile); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        peopleRecords = append(peopleRecords, &people)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return peopleRecords, nil
}
func(peopleRepo *peopleRepository) GetPeople(ctx context.Context, filter, value string) (*models.People, error){

	sqlStatement := fmt.Sprintf("SELECT id ,rut,dv,name,lastname,address,city,commune,email,phone,profile_id FROM %s WHERE %s=%s;", peopleRepo.table,filter, value)
var res models.People
row := peopleRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.Rut, &res.DV,&res.Name, &res.Lastname,&res.Address,&res.City, &res.Commune,&res.Email,&res.Phone,&res.Profile)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}
	return &res, nil
}
func(peopleRepo *peopleRepository) UpdatePeople(ctx context.Context , data *models.People) (*models.People, error){
 
query :=	fmt.Sprintf("UPDATE  %s set rut='%d', dv='%s',name=%s,lastname=%s,address=%s,city=%s,commune=%s,email=%s,phone=%s,profile_id=%d",peopleRepo.table,
			data.Rut, 	
			data.DV,
			data.Name,
			data.Lastname,
			data.Address,
			data.City,
			data.Commune,
			data.Email,
			data.Phone,
			data.Profile,

)
	fmt.Println("Query to execute ", query)
	result, err := peopleRepo.db.Exec(query)
  
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

func(peopleRepo *peopleRepository) 	DeletePeople(ctx context.Context, id string) (error){
	
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", peopleRepo.table,id)

	result, err :=	peopleRepo.db.Exec(query)

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