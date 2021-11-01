package repository

import (
	"context"
	"database/sql"
	"flypack/models"
	"fmt"
)

type CompanyRepository interface {
	GetCompany(ctx context.Context, filter, value string) (*models.Company, error)
	GetAllCompany(ctx context.Context) ([]*models.Company, error)
	CreateCompany(ctx context.Context, 	user *models.Company) (*models.Company, error)
	UpdateCompany(
		ctx context.Context,
		user *models.Company)(*models.Company, error)
	DeleteCompany(ctx context.Context, id string) (error)
}

type companyRepository struct {
	db *sql.DB
	table string
}

func NewCompanyRepository(db *sql.DB) (CompanyRepository, error) {

		return &companyRepository{db: db, table: "companies"}, nil
}

func(companyRepo *companyRepository) CreateCompany(ctx context.Context,data *models.Company) (*models.Company, error){

query :=	fmt.Sprintf("INSERT INTO %s (rut, dv, razon, address, city, commune, people_id) VALUES (%d','%s',%s,%s,%s,%s,%d )", companyRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,
)
	fmt.Println("Query to execute ", query)
	result, err := companyRepo.db.Exec(query)

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
func(companyRepo *companyRepository) GetAllCompany(ctx context.Context) ([]*models.Company, error){


	query := fmt.Sprintf("SELECT  id, rut, dv, razon, address, city, commune, people_id from %s", companyRepo.table)
rows, err := companyRepo.db.Query(query)
    if err != nil {
				fmt.Println("Error db.Query", err.Error())
        return nil, err
    }
    defer rows.Close()

    var companyRecords []*models.Company
    for rows.Next() {
        var company models.Company
        if err := rows.Scan(&company.ID, &company.Rut,&company.DV,
            &company.Razon, &company.Address,&company.City, &company.Commune, &company.LegalPerson); err != nil {
							fmt.Println("Error rows.Scan ", err.Error())
            break
        }
        companyRecords = append(companyRecords, &company)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Error rows.Err ", err.Error())
        return nil, err
    }
    return companyRecords, nil
}
func(companyRepo *companyRepository) GetCompany(ctx context.Context, filter, value string) (*models.Company, error){

	sqlStatement := fmt.Sprintf("SELECT id, rut, dv, razon, address, city, commune, people_id FROM %s WHERE %s=%s;", companyRepo.table,filter, value)
var res models.Company
row := companyRepo.db.QueryRow(sqlStatement)
err := row.Scan(&res.ID, &res.Rut, &res.DV,&res.Razon, &res.Address,&res.City, &res.Commune,&res.LegalPerson)

if err != nil {
	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
  	return nil, nil
	}
		return nil, err
	}
	return &res, nil
}
func(companyRepo *companyRepository) UpdateCompany(ctx context.Context , data *models.Company) (*models.Company, error){

query :=	fmt.Sprintf("UPDATE  %s set rut='%d', dv='%s',razon=%s, address=%s,city=%s,commune=%s,people_id=%d",companyRepo.table,
			data.Rut, 	
			data.DV,
			data.Razon,
			data.Address,
			data.City,
			data.Commune,
			data.LegalPerson,

)
	fmt.Println("Query to execute ", query)
	result, err := companyRepo.db.Exec(query)
  
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
func(companyRepo *companyRepository) 	DeleteCompany(ctx context.Context, id string) (error){
	
	query :=	fmt.Sprintf("DELETE from %s WHERE id='%s'", companyRepo.table,id)

	result, err :=	companyRepo.db.Exec(query)

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