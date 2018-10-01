package model

import (
	"fmt"

	"../config"
	_ "github.com/go-sql-driver/mysql"
)

func FindAll(e *[]Employee) (err error) {
	if err = config.DB.Find(e).Error; err != nil {
		return err
	}
	return nil
}

func Insert(e *Employee) (err error) {
	if err = config.DB.Create(e).Error; err != nil {
		return err
	}
	return nil
}

func FindById(e *Employee, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(e).Error; err != nil {
		return err
	}
	return nil
}

func Update(e *Employee, id string) (err error) {
	employeeTemp := config.DB.First(&e)
	fmt.Println("e temp : ", employeeTemp)
	config.DB.Save(e)
	return nil
}

func Delete(e *Employee, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(e)
	return nil
}
