package main

import (
	"fmt"
	"github/mauljurassicia/nutrifeat/config"
	"github/mauljurassicia/nutrifeat/infrastructure/db"
	"log"
)

func main() {
	viperConfig := config.NewViper()
	var db db.DB
	if !viperConfig.GetBool("app.disable_database") {
		db = config.NewDatabase(viperConfig)
		defer func() {
			if db != nil {
				db.Close()
			}
		}()
	}
	server := config.NewServer(viperConfig)

	fmt.Println("Starting server...")

	config.Bootstrap(config.BootstrapConfig{DB: db, Server: server})

	server.Assets("public")

	log.Fatal(server.Serve(viperConfig.GetString("app.host") + ":" + viperConfig.GetString("app.port")))




}
