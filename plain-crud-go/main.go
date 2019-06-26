package main

import (
	"os"

	"./router"
	"github.com/apsdehal/go-logger"
)

func main() {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}

	log.Debug("Starting App")

	r := router.InitV1()
	r.Run()
}
