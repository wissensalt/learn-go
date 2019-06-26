package dao

import (
	"database/sql"

	"../dto"
	"../global"
	"../util"
	"github.com/google/logger"
)

const (
	QUERY_FIND_ALL_EMPLOYEE   = "SELECT * FROM employee"
	QUERY_FIND_BY_ID_EMPLOYEE = "SELECT * FROM employee WHERE ID = $1"
	QUERY_INSERT_EMPLOYEE     = "INSERT INTO employee (code, name, address, salary) VALUES ($1, $2, $3, $4)"
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

		util.LogInfo(global.LoggerFile, global.Logger, "Received Data : ", tempEmployee)

		result = append(result, tempEmployee)
	}
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
		logger.Warning("No Rows Found")
	case nil:
		result.Id = idR
		result.Code = code
		result.Name = name
		result.Address = address
		result.Salary = salary
	default:
		util.LogError("An Error Occured ", err)
	}
	return result
}

func InsertEmployee(p_EmployeeDTO dto.Employee) (dto.Employee, error) {
	_, err := global.ActiveDB.Query(QUERY_INSERT_EMPLOYEE, p_EmployeeDTO.Code, p_EmployeeDTO.Name, p_EmployeeDTO.Address, p_EmployeeDTO.Salary)
	return p_EmployeeDTO, err
}
