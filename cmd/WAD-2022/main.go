package main

import (
	"WAD-2022/internal/pkg/app"
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	ctx := context.Background()
	log.Println("app start")

	application, err := app.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t create application")

		os.Exit(2)
	}

	err = application.Run(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t run application")

		os.Exit(2)
	}
	log.Println("app terminated")
}
