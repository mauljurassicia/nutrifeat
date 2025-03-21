package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/db"

	"github.com/spf13/viper"
)

func NewDatabase(config *viper.Viper) db.DB {

	databaseConfig := db.DatabaseConfig{
		Host:     config.GetString("database.host"),
		Port:     config.GetInt("database.port"),
		User:     config.GetString("database.user"),
		Password: config.GetString("database.password"),
		Database: config.GetString("database.database"),
	}

	db := &db.SqlGo{}
	db.Open(databaseConfig)

	return db
}
