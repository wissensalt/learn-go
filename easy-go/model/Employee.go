package model

import (
	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	Name string `json:"name"`
	Role string `json:"role"`
}

func (e *Employee) getTableName() string {
	return "employee"
}
