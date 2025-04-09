package patient

import "github/mauljurassicia/nutrifeat/infrastructure/http"

type PatientController struct{}

func NewPatientController() *PatientController {
	return &PatientController{}
}

func (*PatientController) Get(c http.HttpContext) error {
	return nil
}

func (*PatientController) Show(c http.HttpContext) error {
	return nil
}

func (*PatientController) Create(c http.HttpContext) error {
	return nil
}

func (*PatientController) Update(c http.HttpContext) error {
	return nil
}

func (*PatientController) Delete(c http.HttpContext) error {
	return nil
}
