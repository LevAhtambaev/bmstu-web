package main

import (
	"WAD-2022/internal/pkg/app"
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

// @title Comics Store
// @version 1.0
// @description Comics store
// @contact.name API Support
// @contact.url https://vk.com/id250446192
// @contact.email fotchin@mail.ru

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1:8080
// @schemes http https
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	ctx := context.Background()
	log.Println("app start")

	application, err := app.New(ctx)
	if err != nil {
		log.Printf("cant create application: %s", err)

		os.Exit(2)
	}

	err = application.Run()
	if err != nil {
		log.Printf("can`t run application: %s", err)

		os.Exit(2)
	}
	log.Println("app terminated")
}
