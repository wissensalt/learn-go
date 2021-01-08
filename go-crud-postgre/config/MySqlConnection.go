package config

import (
	"bytes"
	"database/sql"
	"os"
	"strconv"
	"time"

	"../global"
	"github.com/joho/godotenv"
)

func MySqlConnect() (*sql.DB, error) {
	err := godotenv.Load()
	var activeDB *sql.DB
	if err != nil {
		global.Logger.Fatal("Error Loading .env variables")
	} else {
		var bufferURLDB bytes.Buffer
		bufferURLDB.WriteString(os.Getenv("db.username"))
		bufferURLDB.WriteString(":")
		bufferURLDB.WriteString(os.Getenv("db.password"))
		bufferURLDB.WriteString("@tcp(")
		bufferURLDB.WriteString(os.Getenv("db.host"))
		bufferURLDB.WriteString(":")
		bufferURLDB.WriteString(os.Getenv("db.port"))
		bufferURLDB.WriteString(")/")
		bufferURLDB.WriteString(os.Getenv("db.name"))

		activeDB, err = sql.Open(os.Getenv("db.instance"), bufferURLDB.String())
		if err != nil {
			global.Logger.Error("Error Connecting to DB")
		} else {
			var maxOpenConnection int
			var maxIdleConnection int
			maxOpenConnection, err = strconv.Atoi(os.Getenv("db.max-open-connection"))
			maxIdleConnection, err = strconv.Atoi(os.Getenv("db.max-idle-connection"))
			activeDB.SetMaxOpenConns(maxOpenConnection)
			activeDB.SetMaxIdleConns(maxIdleConnection)
			activeDB.SetConnMaxLifetime(time.Hour)
		}
	}
	return activeDB, err
}
