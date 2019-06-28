package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"../global"
	"github.com/google/logger"
	"github.com/joho/godotenv"
)

var currentLogFileIncrement int

// InitLogger will open os.File and return pointer of logger.Logger
// appender will roll incremently, the last log must be in highest number
// maximum size for every log is 50 MB
func InitLogger() (*os.File, *logger.Logger) {
	err := godotenv.Load()
	// flag.Parse()

	lf, err := os.OpenFile(DefineAppender(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	loggerObj := logger.Init(os.Getenv("log.name"), false, true, lf)
	return lf, loggerObj
}

// CloseLogger will close open file and logger.Logger reference
func CloseLogger(lf *os.File, logger *logger.Logger) {
	defer lf.Close()
	defer logger.Close()
}

// DefineAppender will get file name to open and conduct rolling
func DefineAppender() string {
	dirContent := ListDirectoryContent(os.Getenv("log.path"))
	lastLog := dirContent[len(dirContent)-1]
	fmt.Println("last log : ", lastLog)

	var logFileToCreate string
	if strings.Contains(lastLog, ".log") {
		maxFileSize, _ := strconv.ParseInt(os.Getenv("log.max-file-size"), 10, 64)
		if CheckFileSizeByPath(lastLog) >= maxFileSize {
			var err error
			if strings.Contains(lastLog, "_") {
				logIncrement := strings.Split(lastLog, "_")
				currentLogFileIncrement, err = strconv.Atoi(logIncrement[1])
				if err != nil {
					panic(err)
				}
			}
			currentLogFileIncrement++
			logFileToCreate = os.Getenv("log.path") + os.Getenv("log.file") + "_" + strconv.Itoa(currentLogFileIncrement)
		} else {
			logFileToCreate = lastLog
		}
	} else {
		logFileToCreate = os.Getenv("log.path") + os.Getenv("log.file")
	}

	return logFileToCreate
}

func RollingLog() {
	if (currentLogFileIncrement <= 9) && (CheckFileSize(global.LoggerFile) >= maxLogSize()) {
		global.Logger.Info("End of Logger")
		CloseLogger(global.LoggerFile, global.Logger)
		InjectLogger()
	}
}

func maxLogSize() int64 {
	maxFileSize, _ := strconv.ParseInt(os.Getenv("log.max-file-size"), 10, 64)
	return maxFileSize
}

func InjectLogger() {
	global.LoggerFile, global.Logger = InitLogger()
}
