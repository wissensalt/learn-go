package util

import (
	"database/sql"

	"github.com/google/logger"
)

func CheckCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			logger.Error("Error counting rows")
		}
	}
	RollingLog()
	return count
}
