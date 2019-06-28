package dao

import (
	"database/sql"

	"../dto"
	"../global"
	"../util"
)

const (
	QUERY_FIND_ALL_EMPLOYEE   = "SELECT * FROM employee"
	QUERY_FIND_BY_ID_EMPLOYEE = "SELECT * FROM employee WHERE ID = $1"
	QUERY_INSERT_EMPLOYEE     = "INSERT INTO employee (code, name, address, salary) VALUES ($1, $2, $3, $4)"
	QUERY_UPDATE_EMPLOYEE     = "UPDATE employee set code=$1, name=$2, address=$3, salary=$4 WHERE id=$1"
	QUERY_DELETE_EMPLOYEE     = "DELETE employee WHERE id=$1"
)

func FindAllEmployee() []dto.Employee {
	var result []dto.Employee

	rows, err := global.ActiveDB.Query(QUERY_FIND_ALL_EMPLOYEE)
	if err != nil {
		global.Logger.Fatal("Error Connect To DB : ", err)
	}

	for rows.Next() {
		var id int
		var code string
		var name string
		var address string
		var salary float64
		err = rows.Scan(&id, &code, &name, &address, &salary)

		var tempEmployee dto.Employee
		tempEmployee.Id = id
		tempEmployee.Code = code
		tempEmployee.Name = name
		tempEmployee.Address = address
		tempEmployee.Salary = salary

		global.Logger.Info("Received Data : ", tempEmployee)

		result = append(result, tempEmployee)
	}
	util.RollingLog()
	return result
}

func FindByIdEmployee(id int) dto.Employee {
	var result dto.Employee
	var idR int
	var code string
	var name string
	var address string
	var salary float64
	row := global.ActiveDB.QueryRow(QUERY_FIND_BY_ID_EMPLOYEE, id)
	switch err := row.Scan(&idR, &code, &name, &address, &salary); err {
	case sql.ErrNoRows:
		global.Logger.Warning("No Rows Found")
	case nil:
		result.Id = idR
		result.Code = code
		result.Name = name
		result.Address = address
		result.Salary = salary
	default:
		global.Logger.Error("An Error Occured ", err)
	}
	util.RollingLog()
	return result
}

func InsertEmployee(p_EmployeeDTO dto.Employee) (dto.Employee, error) {
	tx, err := global.ActiveDB.Begin()
	_, err = global.ActiveDB.Query(QUERY_INSERT_EMPLOYEE, p_EmployeeDTO.Code, p_EmployeeDTO.Name, p_EmployeeDTO.Address, p_EmployeeDTO.Salary)
	if err != nil {
		tx.Rollback()
		global.Logger.Error("Failed Insert Employee")
		util.RollingLog()
	}
	util.CheckErr(tx.Commit())

	return p_EmployeeDTO, err
}

func UpdateEmployee(p_EmployeeDTO dto.Employee) (dto.Employee, error) {
	tx, err := global.ActiveDB.Begin()
	_, err = global.ActiveDB.Query(QUERY_UPDATE_EMPLOYEE, p_EmployeeDTO.Code, p_EmployeeDTO.Name, p_EmployeeDTO.Address, p_EmployeeDTO.Salary)
	if err != nil {
		tx.Rollback()
		global.Logger.Error("Error Update Employee")
		util.RollingLog()
	}
	util.CheckErr(tx.Commit())
	return p_EmployeeDTO, err
}

func DeleteEmployee(p_EmployeeId int) error {
	tx, err := global.ActiveDB.Begin()
	_, err = global.ActiveDB.Query(QUERY_DELETE_EMPLOYEE, p_EmployeeId)
	if err != nil {
		tx.Rollback()
		global.Logger.Error("Error delete Employee")
		util.RollingLog()
	}
	util.CheckErr(tx.Commit())
	return err
}
