package controller

import (
	"../dao"
	"../dto"
	"../util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
}

func GetVisitor(c *gin.Context) {

}

func FindAll(c *gin.Context) {
	employee, error := dao.FindAll()
	if error != nil {
		util.ResponseJSON(c, 400, employee)
	} else {
		util.ResponseJSON(c, 200, employee)
	}
}

func FindById(c *gin.Context) {
	var requestFindById dto.RequestFindById
	c.BindJSON(&requestFindById)
	employee, error := dao.FindById(requestFindById.Id)
	if error != nil {
		util.ResponseJSON(c, 400, employee)
	} else {
		util.ResponseJSON(c, 200, employee)
	}
}

func Insert(c *gin.Context) {
	var requestInsertEmployee dto.RequestInsertEmployee
	c.BindJSON(&requestInsertEmployee)
	error := dao.Insert(&requestInsertEmployee)
	if error != nil {

	}
}
