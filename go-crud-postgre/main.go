package main

import (
	"database/sql"

	"./config"
	"./global"
	"./router"
	"./util"
	"github.com/google/logger"
)

func main() {
	util.InjectLogger()

	activeDB, err := config.Connect()
	if err != nil {
		logger.Error("An Error Occured ", err)
	} else {
		InjectDB(activeDB, global.Logger)
	}

	util.LogInfo("#INITATE ROUTER")
	router.InitRouter()

	util.CloseLogger(global.LoggerFile, global.Logger)
}

func InjectDB(db *sql.DB, logger *logger.Logger) {
	if db == nil {
		logger.Error("Sql DB is NIL, Init new Connection")
		db, _ = config.Connect()
	}
	global.ActiveDB = db
}
