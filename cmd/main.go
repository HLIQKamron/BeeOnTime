package main

import (
	"fmt"
	"log"

	"github.com/BeeOntime/api"
	"github.com/BeeOntime/config"
	"github.com/BeeOntime/pkg/db"
	"github.com/BeeOntime/storage"
)

func main() {
	cfg := config.Load()

	db, err := db.New(cfg)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("db", db)
	}

	router := api.SetUpAPI(cfg, storage.New(db, cfg))

	if err := router.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", err)
	}
}
