package config

import (
	"bytes"
	"database/sql"
	"os"

	"github.com/apsdehal/go-logger"
	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection() (*sql.DB, error) {
	log, err := logger.New("test", 1, os.Stdout)

	var bufferURLDB bytes.Buffer
	bufferURLDB.WriteString(DB_USERNAME)
	bufferURLDB.WriteString(":")
	bufferURLDB.WriteString(DB_PASSWORD)
	bufferURLDB.WriteString("@tcp(")
	bufferURLDB.WriteString(DB_IP)
	bufferURLDB.WriteString(":")
	bufferURLDB.WriteString(DB_PORT)
	bufferURLDB.WriteString(")/")
	bufferURLDB.WriteString(DB_NAME)

	db, err := sql.Open("mysql", bufferURLDB.String())
	if err != nil {
		log.Error("Error Connecting to DB")
	}

	err = db.Ping()
	if err != nil {
		log.Debug("Successfully Connected")
	}
	return db, err
}
