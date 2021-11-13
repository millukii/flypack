package repository

import (
	"context"
	"database/sql"
	"flypack/models"

)

type ShipmentPackagesRepository interface {
	GetShipmentPackages(ctx context.Context, filter, value string) (*models.ShipmentPackage, error)
	GetAllShipmentPackages(ctx context.Context) ([]*models.ShipmentPackage, error)
	CreateShipmentPackages(ctx context.Context, 	user *models.ShipmentPackage) (*models.ShipmentPackage, error)
	UpdateShipmentPackages(
		ctx context.Context,
		user *models.ShipmentPackage)(*models.ShipmentPackage, error)
	DeleteShipmentPackages(ctx context.Context, id string) (error)
}

type shipmentPackagesRepository struct {
	db *sql.DB
	table string
}

func NewShipmentPackagesRepository(db *sql.DB) (ShipmentPackagesRepository, error) {

		return &shipmentPackagesRepository{db: db, table: "companies"}, nil
}

func(shipmentPackagesRepo *shipmentPackagesRepository) CreateShipmentPackages(ctx context.Context,data *models.ShipmentPackage) (*models.ShipmentPackage, error){
/* 
query :=	fmt.Sprintf("INSERT INTO %s (rut, dv, razon, address, city, commune, people_id) VALUES (%d','%s',%s,%s,%s,%s,%d )", shipmentPackagesRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,
)
	fmt.Println("Query to execute ", query)
	result, err := shipmentPackagesRepo.db.Exec(query)

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
func(shipmentPackagesRepo *shipmentPackagesRepository) GetAllShipmentPackages(ctx context.Context) ([]*models.ShipmentPackage, error){

/* 
	query := fmt.Sprintf("SELECT  id, rut, dv, razon, address, city, commune, people_id from %s", shipmentPackagesRepo.table)
rows, err := shipmentPackagesRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var shipmentPackagesRecords []*models.ShipmentPackages
    for rows.Next() {
        var shipmentPackages models.ShipmentPackages
        if err := rows.Scan(&shipmentPackages.ID, &shipmentPackages.Rut,&shipmentPackages.DV,
            &shipmentPackages.Razon, &shipmentPackages.Address,&shipmentPackages.City, &shipmentPackages.Commune, &shipmentPackages.LegalPerson); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        shipmentPackagesRecords = append(shipmentPackagesRecords, &shipmentPackages)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    } */
    return nil, nil
}
func(shipmentPackagesRepo *shipmentPackagesRepository) GetShipmentPackages(ctx context.Context, filter, value string) (*models.ShipmentPackage, error){
/* 
	sqlStatement := fmt.Sprintf("SELECT id, rut, dv, razon, address, city, commune, people_id FROM %s WHERE %s=%s;", shipmentPackagesRepo.table,filter, value)
var res models.ShipmentPackages
row := shipmentPackagesRepo.db.QueryRow(sqlStatement)
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

func(shipmentPackagesRepo *shipmentPackagesRepository) UpdateShipmentPackages(ctx context.Context , data *models.ShipmentPackage) (*models.ShipmentPackage, error){
/*
query :=	fmt.Sprintf("UPDATE  %s set rut='%d', dv='%s',razon=%s, address=%s,city=%s,commune=%s,people_id=%d",shipmentPackagesRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,

)
	fmt.Println("Query to execute ", query)
	result, err := shipmentPackagesRepo.db.Exec(query)
  
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
func(shipmentPackagesRepo *shipmentPackagesRepository) 	DeleteShipmentPackages(ctx context.Context, id string) (error){
	/* 
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", shipmentPackagesRepo.table,id)

	result, err :=	shipmentPackagesRepo.db.Exec(query)

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