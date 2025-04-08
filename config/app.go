package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/db"
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	route "github/mauljurassicia/nutrifeat/presentation/controller/config"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/auth"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/dashboard"
	"github/mauljurassicia/nutrifeat/presentation/controller/web/home"
)

type BootstrapConfig struct {
	DB     db.DB
	Server http.HttpApp
}

func Bootstrap(config BootstrapConfig) {
	homeController := home.NewHomeController()
	authController := auth.NewAuthController()
	dashboardController := dashboard.NewDashboardController()

	routeConfig := route.RouteConfig{App: config.Server,
		HomeController:      homeController,
		DashboardController: dashboardController,
		AuthController:      authController}
	route.RegisterRoutes(&routeConfig)

}
