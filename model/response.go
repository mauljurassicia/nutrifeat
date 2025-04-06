package model


import  (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Methode string      `json:"methode"`
	Path    string      `json:"path"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Methode string      `json:"methode"`
	Path    string      `json:"path"`
	Version string      `json:"version"`
	Errors  interface{} `json:"errors"`
}



func NewErrorResponse(c http.HttpContext, status int, message string, errors interface{}) ErrorResponse {
	return ErrorResponse{Status: status, Message: message, Errors: errors, Methode: c.Methode(), Path: c.Path(), Version: "1.0.0"}
}

func NewResponse(c http.HttpContext, status int, message string, data interface{}) Response {
	return Response{Status: status, Message: message, Data: data, Methode: c.Methode(), Path: c.Path(), Version: "1.0.0"}
}