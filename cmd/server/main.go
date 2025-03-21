package main

import (
	"fmt"
	"github/mauljurassicia/nutrifeat/config"
	"log"
)

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	server := config.NewServer(viperConfig)

	fmt.Println("Starting server...")

	config.Bootstrap(config.BootstrapConfig{DB: db, Server: server})

	defer db.Close()

	log.Fatal(server.Serve(viperConfig.GetString("app.host") + ":" + viperConfig.GetString("app.port")))




}
