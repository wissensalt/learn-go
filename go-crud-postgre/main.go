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
	global.Logger.Info("#LOGGER INITATED")

	activeDB, err := config.PostgreSQLConnect()
	if err != nil {
		logger.Error("An Error Occured ", err)
	} else {
		InjectDB(activeDB, global.Logger)
	}

	router.InitRouter()
	util.CloseLogger(global.LoggerFile, global.Logger)
}

func InjectDB(db *sql.DB, logger *logger.Logger) {
	if db == nil {
		logger.Error("Sql DB is NIL, Init new Connection")
		db, _ = config.PostgreSQLConnect()
	}
	global.ActiveDB = db
	global.Logger.Info("#DB CONNECTION ESTABLISHED")
}
