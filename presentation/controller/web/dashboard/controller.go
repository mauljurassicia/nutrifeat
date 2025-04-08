package dashboard

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/presentation/view/pages/dashboard"
)

type DashboardController struct {
	//
}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

func (controller *DashboardController) GetDashboard(c http.HttpContext) error {
	return c.Render(dashboard.Dashboard())
}