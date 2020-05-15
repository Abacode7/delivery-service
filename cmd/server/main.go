package main

import (
	"github.com/Abacode7/delivery-service/pkg/http/rest"
	"github.com/Abacode7/delivery-service/pkg/service/user"
	"github.com/Abacode7/delivery-service/pkg/storage/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func runApp() error {
	// Load environment variable into the system
	err := godotenv.Load()
	if err != nil{
		return err
	}

	// Get environment variables
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	// Initialise and inject repository and services (used in handlers)
	userRepository := mysql.NewStorage(username, password, dbHost, dbName)
	userService := user.NewService(&userRepository)
	handler := rest.Handlers(&userService)

	// Initialise and start server
	srv := &http.Server{Addr: ":8080", Handler: handler, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second}
	er := srv.ListenAndServe()
	if er != nil{
		return er
	}
	return nil
}

func main()  {
	if err := runApp(); err != nil{
		log.Fatal(err)
	}
}

