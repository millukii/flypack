package repository

import (
	"context"
	"database/sql"
	"flypack/models"

)

type ManifiestRepository interface {
	GetManifiest(ctx context.Context, filter, value string) (*models.Manifiest, error)
	GetAllManifiest(ctx context.Context) ([]*models.Manifiest, error)
	CreateManifiest(ctx context.Context, 	user *models.Manifiest) (*models.Manifiest, error)
	UpdateManifiest(
		ctx context.Context,
		user *models.Manifiest)(*models.Manifiest, error)
	DeleteManifiest(ctx context.Context, id string) (error)
}

type manifiestRepository struct {
	db *sql.DB
	table string
}

func NewManifiestRepository(db *sql.DB) (ManifiestRepository, error) {

		return &manifiestRepository{db: db, table: "companies"}, nil
}

func(manifiestRepo *manifiestRepository) CreateManifiest(ctx context.Context,data *models.Manifiest) (*models.Manifiest, error){
/* 
query :=	fmt.Sprintf("INSERT INTO %s (rut, dv, razon, address, city, commune, people_id) VALUES (%d','%s',%s,%s,%s,%s,%d )", manifiestRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,
)
	fmt.Println("Query to execute ", query)
	result, err := manifiestRepo.db.Exec(query)

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
func(manifiestRepo *manifiestRepository) GetAllManifiest(ctx context.Context) ([]*models.Manifiest, error){

/* 
	query := fmt.Sprintf("SELECT  id, rut, dv, razon, address, city, commune, people_id from %s", manifiestRepo.table)
rows, err := manifiestRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var manifiestRecords []*models.Manifiest
    for rows.Next() {
        var manifiest models.Manifiest
        if err := rows.Scan(&manifiest.ID, &manifiest.Rut,&manifiest.DV,
            &manifiest.Razon, &manifiest.Address,&manifiest.City, &manifiest.Commune, &manifiest.LegalPerson); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        manifiestRecords = append(manifiestRecords, &manifiest)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    } */
    return nil, nil
}
func(manifiestRepo *manifiestRepository) GetManifiest(ctx context.Context, filter, value string) (*models.Manifiest, error){
/* 
	sqlStatement := fmt.Sprintf("SELECT id, rut, dv, razon, address, city, commune, people_id FROM %s WHERE %s=%s;", manifiestRepo.table,filter, value)
var res models.Manifiest
row := manifiestRepo.db.QueryRow(sqlStatement)
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

func(manifiestRepo *manifiestRepository) UpdateManifiest(ctx context.Context , data *models.Manifiest) (*models.Manifiest, error){
/*
query :=	fmt.Sprintf("UPDATE  %s set rut='%d', dv='%s',razon=%s, address=%s,city=%s,commune=%s,people_id=%d",manifiestRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,

)
	fmt.Println("Query to execute ", query)
	result, err := manifiestRepo.db.Exec(query)
  
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
func(manifiestRepo *manifiestRepository) 	DeleteManifiest(ctx context.Context, id string) (error){
	/* 
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", manifiestRepo.table,id)

	result, err :=	manifiestRepo.db.Exec(query)

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