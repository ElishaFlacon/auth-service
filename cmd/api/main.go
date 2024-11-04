package main

import (
	"log"

	"github.com/ElishaFlacon/shelly/internal/app"
)

const (
	errMsgFailedToInitApp = "failed to init app: %s"
	errMsgFailedToRunApp  = "failed to run app: %s"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf(errMsgFailedToInitApp, err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf(errMsgFailedToRunApp, err.Error())
	}
}
