package home

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/presentation/view/pages"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (controller *HomeController) GetHome(c http.HttpContext) error {
	return c.Render(pages.Home())
}

func (controller *HomeController) GetUserName(c http.HttpContext) error {
	username := c.Query("username")

	return c.JSON(map[string]string{"username": username})
}

func (controller *HomeController) GetAllParams(c http.HttpContext) error {
	return c.JSON(c.AllParams())
}
