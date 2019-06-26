package util

import "github.com/google/logger"

func CheckErr(err error, logger *logger.Logger) {
	if err != nil {
		logger.Fatal("An Error occured : ", err)
	}
}
