package http

import (
	"net/http"
	"github.com/a-h/templ"
)




type HttpApp interface {
	Get(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Post(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Options(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Head(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Patch(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Delete(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Put(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Serve(addr string) error
	Assets(path string)
	Group(prefix string, middleware ...MiddlewareFuncWithContext) HttpGroup
	Use(middleware ...MiddlewareFuncWithContext)
}

type HttpGroup interface {
	Get(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Post(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Options(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Head(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Patch(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Delete(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
	Put(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error
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
	Render(view templ.Component ) error

}


type MiddlewareFuncWithContext func(HandlerFuncWithContext) HandlerFuncWithContext


type HandlerFuncWithContext func(c HttpContext) error

type NewErrorHandler func (err error) http.HandlerFunc
