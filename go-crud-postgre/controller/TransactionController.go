package controller

import (
	"net/http"
	"strconv"

	"../dao"
	"../dto"
	"../global"
	"../util"
	"../validator"
	"github.com/gin-gonic/gin"
)

func InsertEmployee(p_Context *gin.Context) {
	var payload dto.Employee
	var err error
	var resultInsert dto.Employee

	err = p_Context.BindJSON(&payload)

	if err != nil {
		util.BuildResponseBadRequest(p_Context, strconv.Itoa(http.StatusBadRequest), "Bad Request")
	} else {
		if validator.ValidateEmployee(payload) {
			resultInsert, err = dao.InsertEmployee(payload)

			if err != nil {
				global.Logger.Error("Error Insert Employee", err)
				util.RollingLog()
				util.BuildResponseInternalServerError(p_Context, strconv.Itoa(http.StatusInternalServerError), "Error Insert Employee")
			} else {
				util.BuildResponseOKWithContent(p_Context, strconv.Itoa(http.StatusOK), "Success", resultInsert)
			}
		} else {
			util.BuildResponseBadRequest(p_Context, strconv.Itoa(http.StatusBadRequest), "Bad Request")
		}
	}
}

func UpdateEmployee(p_Context *gin.Context) {
	var payload dto.Employee
	var err error
	err = p_Context.BindJSON(&payload)

	if err != nil {
		util.BuildResponseBadRequest(p_Context, strconv.Itoa(http.StatusBadRequest), "Bad Request")
	} else {
		payload, err = dao.UpdateEmployee(payload)
		if err != nil {
			util.BuildResponseInternalServerError(p_Context, strconv.Itoa(http.StatusInternalServerError), "Error update employee")
		} else {
			util.BuildResponseOKWithContent(p_Context, "200", "OK", payload)
		}
	}
}

func DeleteEmployee(p_Context *gin.Context) {
	var payload dto.RequestFindById
	var err error
	err = p_Context.BindJSON(&payload)

	if err != nil {
		util.BuildResponseBadRequest(p_Context, strconv.Itoa(http.StatusBadRequest), "Bad Request")
	} else {
		err = dao.DeleteEmployee(payload.Id)
		if err != nil {
			util.BuildResponseInternalServerError(p_Context, strconv.Itoa(http.StatusInternalServerError), "Error Delete Employee")
		} else {
			util.BuildResponseOK(p_Context, strconv.Itoa(http.StatusOK), "Success Delete Employee")
		}
	}
}
