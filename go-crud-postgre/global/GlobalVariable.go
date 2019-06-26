package global

import (
	"database/sql"
	"os"

	"github.com/google/logger"
)

var Logger *logger.Logger
var LoggerFile *os.File
var ActiveDB *sql.DB
