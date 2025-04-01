package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/auth"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/home"
)

type RouteConfig struct {
	App http.HttpApp
	HomeController *home.HomeController
	AuthController *auth.AuthController
}

func RegisterRoutes(c *RouteConfig) {
	c.App.Get("/", c.HomeController.GetHome)
	c.App.Get("/user", c.HomeController.GetUserName)
	c.App.Get("/params/{username}/{password}", c.HomeController.GetAllParams)
	c.App.Get("/login", c.AuthController.GetLogin)
	c.App.Get("/register", c.AuthController.GetRegister)
}

