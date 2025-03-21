package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"

	"github.com/spf13/viper"
)

func NewServer(config *viper.Viper) http.HttpApp {

	app := http.NewStd()

	return app
}