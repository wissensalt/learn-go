package controller

import (
	"github.com/gin-gonic/gin"

	"../dao"
	"../dto"
	"../util"
)

func FindAll(p_Context *gin.Context) {
	allEmployee := dao.FindAllEmployee()
	var responseData dto.ResponseData
	responseData.ResponseCode = "200"
	responseData.ResponseMsg = "Success"
	util.BuildResponseDataWithContent(p_Context, responseData, allEmployee)
}

func FindById(p_Context *gin.Context) {
	var requestFindByID dto.RequestFindById
	p_Context.BindJSON(&requestFindByID)

	var responseData dto.ResponseData
	responseData.ResponseCode = "200"
	responseData.ResponseMsg = "Success"
	employee := dao.FindByIdEmployee(requestFindByID.Id)
	util.BuildResponseDataWithContent(p_Context, responseData, employee)
}
