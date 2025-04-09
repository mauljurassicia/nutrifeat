package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/auth"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/dashboard"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/home"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/patient"
)

type RouteConfig struct {
	App                 http.HttpApp
	HomeController      *home.HomeController
	AuthController      *auth.AuthController
	DashboardController *dashboard.DashboardController
	PatientController   *patient.PatientController
}

func RegisterRoutes(c *RouteConfig) {
	c.App.Get("/", c.HomeController.GetHome)
	c.App.Get("/user", c.HomeController.GetUserName)
	c.App.Get("/params/{username}/{password}", c.HomeController.GetAllParams)
	c.App.Get("/login", c.AuthController.GetLogin)
	c.App.Get("/register", c.AuthController.GetRegister)
	c.App.Post("/login", c.AuthController.PostLogin)
	c.App.Post("/register", c.AuthController.PostRegister)
	c.App.Get("/dashboard", c.DashboardController.GetDashboard)


	// patient routes
	patientRoute := c.App.Group("/patients")
	patientRoute.Get("/", c.PatientController.Get)
	patientRoute.Get("/{id}", c.PatientController.Show)
	patientRoute.Post("/", c.PatientController.Create)
	patientRoute.Delete("/{id}", c.PatientController.Delete)
	patientRoute.Patch("/{id}", c.PatientController.Update)


	// room routes


	// food routes


	// recipe routes

	//
}
