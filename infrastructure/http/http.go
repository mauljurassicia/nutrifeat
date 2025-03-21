package http

import (
	"net/http"

)




type HttpApp interface {
	Get(path string, handle HandlerFuncWithContext) error
	Post(path string, handle HandlerFuncWithContext) error
	Options(path string, handle HandlerFuncWithContext) error
	Head(path string, handle HandlerFuncWithContext) error
	Patch(path string, handle HandlerFuncWithContext) error
	Delete(path string, handle HandlerFuncWithContext) error
	Put(path string, handle HandlerFuncWithContext) error
	Serve(addr string) error
}

type HttpContext interface {
	Accepts(offers ...string) string
	AllParams() map[string]string
	Send(body []byte) error
	SendStatus(status int) error
	Params(key string, defaultValue ...string) string
	Query(key string, defaultValue ...string) string
	Protocol() string
	JSON(body interface{}, ctype ...string) error
	BodyParser(body interface{}) error
	Render(template string, data interface{}, layout ...string) error
}

type HandlerFuncWithContext func(c HttpContext) error

type NewErrorHandler func (err error) http.HandlerFunc







