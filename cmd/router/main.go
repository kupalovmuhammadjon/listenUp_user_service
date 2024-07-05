package main

import (
	"log"
	"user_service/api"
	"user_service/config"
	"user_service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	router := api.NewRouter(db)
	router.Run(cfg.HTTP_PORT)
}
