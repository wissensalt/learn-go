package controller

import (
	"fmt"

	"../model"
	"../util"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var employee []model.Employee
	err := model.FindAll(&employee)
	if err != nil {
		util.ResponseJSON(c, 404, employee)
	} else {
		fmt.Println("DATA : ", employee)
		util.ResponseJSON(c, 200, employee)
	}
}

func Insert(c *gin.Context) {
	var employee model.Employee
	c.BindJSON(&employee)
	err := model.Insert(&employee)
	if err != nil {
		util.ResponseJSON(c, 404, employee)
	} else {
		util.ResponseJSON(c, 200, employee)
	}
}

func FindById(c *gin.Context) {
	var employee model.Employee
	id := c.Params.ByName("id")
	fmt.Println("Found ID ", id)
	err := model.FindById(&employee, id)
	if err != nil {
		util.ResponseJSON(c, 404, employee)
	} else {
		util.ResponseJSON(c, 200, employee)
	}
}

func Update(c *gin.Context) {
	var employee model.Employee
	id := c.Params.ByName("id")
	err := model.Update(&employee, id)
	if err != nil {
		util.ResponseJSON(c, 404, employee)
	}
	c.BindJSON(&employee)
	err = model.Update(&employee, id)
	if err != nil {
		util.ResponseJSON(c, 404, employee)
	} else {
		util.ResponseJSON(c, 200, employee)
	}
}

func Delete(c *gin.Context) {
	var employee model.Employee
	id := c.Params.ByName("id")
	err := model.Delete(&employee, id)
	if err != nil {
		util.ResponseJSON(c, 404, employee)
	} else {
		util.ResponseJSON(c, 200, employee)
	}
}
