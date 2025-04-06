package auth

import (
	"fmt"

	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/model"
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
	var registerRequest model.RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		return c.Status(http.BadRequest).JSON(model.NewErrorResponse(c, http.BadRequest, model.ErrParsing, fmt.Sprintf("Error parsing request body: %v", err)))
	}

	errors := model.ValidateStruct(registerRequest)
	if errors != nil {
		return c.Status(http.BadRequest).JSON(model.NewErrorResponse(c, http.BadRequest, model.ErrValidation, errors))
	}

	// TODO: Implement user registration logic here

	return c.Status(http.OK).JSON(model.NewResponse(c, http.OK, "User registered successfully", nil))
}
