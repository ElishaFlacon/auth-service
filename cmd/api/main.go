package main

import (
	"context"
	"log"

	"github.com/ElishaFlacon/auth-service/internal/app"
)

const (
	errMsgFailedToInitApp = "failed to init app: %s"
	errMsgFailedToRunApp  = "failed to run app: %s"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf(errMsgFailedToInitApp, err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf(errMsgFailedToRunApp, err.Error())
	}
}
