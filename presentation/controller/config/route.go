package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/home"
)

type RouteConfig struct {
	App http.HttpApp
	HomeController *home.HomeController
}

func RegisterRoutes(c *RouteConfig) {
	c.App.Get("/", c.HomeController.GetHome)
}

