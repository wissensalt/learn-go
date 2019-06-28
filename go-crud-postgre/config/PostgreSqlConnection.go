package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"../util"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PostgreSQLConnect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env vars")
	}
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("db.host"),
		os.Getenv("db.port"),
		os.Getenv("db.username"),
		os.Getenv("db.password"),
		os.Getenv("db.name"))
	activeDB, err := sql.Open("postgres", dbinfo)
	maxOpenConnection, err := strconv.Atoi(os.Getenv("db.max-open-connection"))
	maxIdleConnection, err := strconv.Atoi(os.Getenv("db.max-idle-connection"))
	// connMaxLifeTime, err := strconv.Atoi(os.Getenv("db.con-max-life-time"))

	activeDB.SetMaxOpenConns(maxOpenConnection)
	activeDB.SetMaxIdleConns(maxIdleConnection)
	activeDB.SetConnMaxLifetime(time.Hour)
	util.CheckErr(err)
	return activeDB, err
}
