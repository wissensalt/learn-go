package main

import (
	"os"

	"./config"
	"./model"
	"./router"
	"github.com/apsdehal/go-logger"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}
	config.DB, err = gorm.Open("mysql", "root:P@ssw0rd@tcp(127.0.0.1:3306)/EmployeeDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Debug("Cannot Connect To DB")
	} else {
		log.Debug("Succesfully Connect To DB")
	}
	defer config.DB.Close()

	log.Debug("Migrating DB")
	config.DB.AutoMigrate(&model.Employee{})

	log.Debug("Setting Routers")
	r := router.SetupRouter()

	log.Debug("Application Running")
	r.Run()
}
