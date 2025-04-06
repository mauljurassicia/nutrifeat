package usecase

import (
	"github/mauljurassicia/nutrifeat/infrastructure/db/repository"
	"github/mauljurassicia/nutrifeat/model"
)

type AuthUsecase interface {
	Login(request model.LoginRequest) error
	Register(request model.RegisterRequest) error
}

func NewAuthUsecase(authRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{authRepo}
}

type authUsecase struct {
	authRepo repository.UserRepository
}

func (a *authUsecase) Login(request model.LoginRequest) error {
	return nil
}

func (a *authUsecase) Register(request model.RegisterRequest) error {
	return nil
}