package dao

import (
	"database/sql"
	"fmt"
	"log"

	"../config"
	"../dto"
	"../model"
)

const (
	QUERY_FIND_ALL_EMPLOYEE = "SELECT * FROM employee;"
	QUERY_FIND_BY_ID        = "SELECT * FROM employee WHERE id=?"
	QUERY_INSERT            = "INSERT INTO employee (name, role) VALUES (?, ?)"
)

func FindAll() ([]model.Employee, error) {
	db, _ := config.CreateConnection()
	rows, err := db.Query(QUERY_FIND_ALL_EMPLOYEE)
	if err != nil {
		log.Fatal("Error ", err)
	}
	defer db.Close()
	defer rows.Close()
	employee := model.Employee{}
	result := []model.Employee{}
	for rows.Next() {
		tempEmployee := new(model.Employee)
		err := rows.Scan(&tempEmployee.Id, &tempEmployee.Name, &tempEmployee.Role)
		if err != nil {
			log.Fatal(err)
		}
		employee.Id = tempEmployee.Id
		employee.Name = tempEmployee.Name
		employee.Role = tempEmployee.Role
		result = append(result, employee)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result, err
}

func FindById(id int) (model.Employee, error) {
	result := model.Employee{}
	db, _ := config.CreateConnection()
	row, err := db.Query(QUERY_FIND_BY_ID, id)
	templEmployee := new(model.Employee)
	for row.Next() {
		err := row.Scan(&templEmployee.Id, &templEmployee.Name, &templEmployee.Role)
		if err == sql.ErrNoRows {
			log.Fatal("No Rows Found")
		}
		if err != nil {
			panic("Error Find By Id")
		} else {
			result.Id = templEmployee.Id
			result.Name = templEmployee.Name
			result.Role = templEmployee.Role
		}
	}
	defer row.Close()
	defer db.Close()

	return result, err
}

func Insert(employee *dto.RequestInsertEmployee) error {
	db, _ := config.CreateConnection()
	statement, err := db.Prepare(QUERY_INSERT)
	if err != nil {

	} else {
		_, err := statement.Exec("Fulan", "Staff")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Success Insert Employee")
		}
	}
	return err
}
