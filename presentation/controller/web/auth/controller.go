package auth

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	login "github/mauljurassicia/nutrifeat/presentation/view/pages/login"
	register "github/mauljurassicia/nutrifeat/presentation/view/pages/register"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) GetLogin(c http.HttpContext) error {
	return c.Render(login.Login())
}

func (controller *AuthController) GetRegister(c http.HttpContext) error {
	return c.Render(register.Register())
}

func (controller *AuthController) PostLogin(c http.HttpContext) error {
	return nil
}

func (controller *AuthController) PostRegister(c http.HttpContext) error {
	return nil
}	
