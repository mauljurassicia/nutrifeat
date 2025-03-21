package home

import "github/mauljurassicia/nutrifeat/infrastructure/http"

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (controller *HomeController) GetHome(c http.HttpContext) error {
	return c.Send([]byte("Hello World!"))
}
