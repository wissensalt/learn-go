package util

import (
	"database/sql"
)

func CheckCount(rows *sql.Rows) (count int) {
	lf, logger := InitLogger()
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			logger.Error("Error counting rows")
		}
	}
	CloseLogger(lf, logger)
	return count
}
