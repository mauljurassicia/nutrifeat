package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/db"
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/home"
	route "github/mauljurassicia/nutrifeat/presentation/controller/config"
)

type BootstrapConfig struct {
	DB db.DB
	Server http.HttpApp
}

func Bootstrap(config BootstrapConfig) {
	homeController := home.NewHomeController()

	routeConfig := route.RouteConfig{App: config.Server, HomeController: homeController}
	route.RegisterRoutes(&routeConfig)



}
