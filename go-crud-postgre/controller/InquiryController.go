package controller

import (
	"github.com/gin-gonic/gin"

	"../dao"
	"../dto"
	"../util"
)

func FindAll(p_Context *gin.Context) {
	allEmployee := dao.FindAllEmployee()
	util.BuildResponseOKWithContent(p_Context, "200", "Success", allEmployee)
}

func FindById(p_Context *gin.Context) {
	var requestFindByID dto.RequestFindById
	p_Context.BindJSON(&requestFindByID)

	employee := dao.FindByIdEmployee(requestFindByID.Id)
	util.BuildResponseOKWithContent(p_Context, "200", "Success", employee)
}
