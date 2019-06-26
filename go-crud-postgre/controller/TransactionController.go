package controller

import (
	"net/http"
	"strconv"

	"../dao"
	"../dto"
	"../util"
	"github.com/gin-gonic/gin"
)

func InsertEmployee(p_Context *gin.Context) {
	var payload dto.Employee
	var responseData dto.ResponseData
	err := p_Context.BindJSON(&payload)

	if err != nil {
		responseData.ResponseCode = strconv.Itoa(http.StatusBadRequest)
		responseData.ResponseMsg = "Bad Request"
	}

	resultInsert, err := dao.InsertEmployee(payload)

	if err != nil {
		util.LogError("Error Insert Employee", err)
	} else {
		responseData.ResponseCode = strconv.Itoa(http.StatusOK)
		responseData.ResponseMsg = "Success"
	}
	util.BuildResponseDataWithContent(p_Context, responseData, resultInsert)
}
