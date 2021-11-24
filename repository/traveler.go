package repository

import (
	"context"
	"database/sql"
	"flypack/models"

)

type TravelerRepository interface {
	GetTraveler(ctx context.Context, filter, value string) (*models.Traveler, error)
	GetAllTraveler(ctx context.Context) ([]*models.Traveler, error)
	CreateTraveler(ctx context.Context, 	user *models.Traveler) (*models.Traveler, error)
	UpdateTraveler(
		ctx context.Context,
		user *models.Traveler)(*models.Traveler, error)
	DeleteTraveler(ctx context.Context, id string) (error)
}

type travelerRepository struct {
	db *sql.DB
	table string
}

func NewTravelerRepository(db *sql.DB) (TravelerRepository, error) {

		return &travelerRepository{db: db, table: "companies"}, nil
}

func(travelerRepo *travelerRepository) CreateTraveler(ctx context.Context,data *models.Traveler) (*models.Traveler, error){
/* 
query :=	fmt.Sprintf("INSERT INTO %s (rut, dv, razon, address, city, commune, people_id) VALUES (%d','%s',%s,%s,%s,%s,%d )", travelerRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,
)
	fmt.Println("Query to execute ", query)
	result, err := travelerRepo.db.Exec(query)

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
func(travelerRepo *travelerRepository) GetAllTraveler(ctx context.Context) ([]*models.Traveler, error){

/* 
	query := fmt.Sprintf("SELECT  id, rut, dv, razon, address, city, commune, people_id from %s", travelerRepo.table)
rows, err := travelerRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var travelerRecords []*models.Traveler
    for rows.Next() {
        var traveler models.Traveler
        if err := rows.Scan(&traveler.ID, &traveler.Rut,&traveler.DV,
            &traveler.Razon, &traveler.Address,&traveler.City, &traveler.Commune, &traveler.LegalPerson); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        travelerRecords = append(travelerRecords, &traveler)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    } */
    return nil, nil
}
func(travelerRepo *travelerRepository) GetTraveler(ctx context.Context, filter, value string) (*models.Traveler, error){
/* 
	sqlStatement := fmt.Sprintf("SELECT id, rut, dv, razon, address, city, commune, people_id FROM %s WHERE %s=%s;", travelerRepo.table,filter, value)
var res models.Traveler
row := travelerRepo.db.QueryRow(sqlStatement)
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

func(travelerRepo *travelerRepository) UpdateTraveler(ctx context.Context , data *models.Traveler) (*models.Traveler, error){
/*
query :=	fmt.Sprintf("UPDATE  %s set rut='%d', dv='%s',razon=%s, address=%s,city=%s,commune=%s,people_id=%d",travelerRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,

)
	fmt.Println("Query to execute ", query)
	result, err := travelerRepo.db.Exec(query)
  
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
func(travelerRepo *travelerRepository) 	DeleteTraveler(ctx context.Context, id string) (error){
	/* 
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", travelerRepo.table,id)

	result, err :=	travelerRepo.db.Exec(query)

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