package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github/mauljurassicia/nutrifeat/public"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

type App struct {
	router       *mux.Router
	errorHandler NewErrorHandler
	middleware   []MiddlewareFuncWithContext
}

func NewStd() *App {
	app := &App{
		router: mux.NewRouter(),
		errorHandler: func(err error) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(err.Error()))
			}
		},
		middleware: []MiddlewareFuncWithContext{},
	}
	app.router.Use(app.recoverMiddleware)
	return app
}

func (app *App) Use(middleware ...MiddlewareFuncWithContext) {
	app.middleware = append(app.middleware, middleware...)
}

func (app *App) Group(prefix string, middleware ...MiddlewareFuncWithContext) HttpGroup {
	return &group{
		prefix:     prefix,
		app:        app,
		middleware: middleware,
	}
}

type group struct {
	prefix     string
	app        *App
	middleware []MiddlewareFuncWithContext
}

func (g *group) Get(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodGet, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

func (g *group) Post(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodPost, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

func (g *group) Put(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodPut, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

func (g *group) Delete(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodDelete, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

func (g *group) Patch(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodPatch, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

func (g *group) Options(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodOptions, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

func (g *group) Head(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return g.app.registerRoute(http.MethodHead, g.prefix+path, handle, append(g.middleware, middleware...)...)
}

type StdContext struct {
	Context        context.Context // Embed the context.Context
	responseWriter http.ResponseWriter
	request        *http.Request
	params         map[string]string // Store path parameters
}

func (c *StdContext) Accepts(offers ...string) string {
	acceptHeader := c.request.Header.Get("Accept")
	for _, offer := range offers {
		if strings.Contains(acceptHeader, offer) {
			return offer
		}
	}
	return ""
}

func (c *StdContext) AllParams() map[string]string {
	return c.params
}

func (c *StdContext) Send(body []byte) error {
	_, err := c.responseWriter.Write(body)
	return err
}

func (c *StdContext) SendStatus(status int) error {
	c.responseWriter.WriteHeader(status)
	return nil
}

func (c *StdContext) Status(status int) HttpContext {
	c.responseWriter.WriteHeader(status)
	return c
}

func (c *StdContext) Params(key string, defaultValue ...string) string {
	if val, exists := c.params[key]; exists {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func (c *StdContext) Query(key string, defaultValue ...string) string {
	value := c.request.URL.Query().Get(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func (c *StdContext) Protocol() string {
	if c.request.TLS != nil {
		return "https"
	}
	return "http"
}

func (c *StdContext) JSON(body interface{}, ctype ...string) error {
	contentType := "application/json"
	if len(ctype) > 0 {
		contentType = ctype[0]
	}
	c.responseWriter.Header().Set("Content-Type", contentType)
	c.responseWriter.WriteHeader(http.StatusOK)
	return json.NewEncoder(c.responseWriter).Encode(body)
}

func (c *StdContext) Cookies(key string, defaultValue ...string) string {
	cookie, err := c.request.Cookie(key)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return cookie.Value
}

func (c *StdContext) BodyParser(body interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(body)
}

func (c *StdContext) Render(view templ.Component) error {
	// Rendering logic should be implemented here, depending on the template engine used

	ctx := context.WithValue(c.Context, "hx-request", c.request.Header.Get("HX-Request") == "true")
	ctx = context.WithValue(ctx, "path", c.request.URL.Path)

	return view.Render(ctx, c.responseWriter)
}

func (c *StdContext) Methode() string {
	return c.request.Method
}

func (c *StdContext) Path() string {
	return c.request.URL.Path
}

func (c *StdContext) Header(key string, defaultValue ...string) string {
	value := c.request.Header.Get(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

func (app *App) Assets(path string) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	fs := public.GetAssets()

	fmt.Println("Serving assets...")
	fmt.Print(fs)
	app.router.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(fs)))
}

// Serve starts the HTTP server.
func (app *App) Serve(addr string) error {
	return http.ListenAndServe(addr, app.router)
}

func (app *App) stdHandler(h HandlerFuncWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Create a new StdContext.
		c := &StdContext{
			Context:        ctx,
			responseWriter: w,
			request:        r,
			params:         make(map[string]string), // Initialize params
		}

		// Call the handler and handle errors.
		err := h(c)
		if err != nil {
			app.errorHandler(err)(w, r)
			return
		}
	}
}

func (app *App) recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Register HTTP methods dynamically.
func (app *App) registerRoute(method string, path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	handler := handle
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	route := app.router.Methods(method).Path(path)
	route.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract path parameters
		vars := mux.Vars(r)
		params := make(map[string]string)
		for key, value := range vars {
			params[key] = value
		}

		// Create a new StdContext.
		ctx := r.Context()
		c := &StdContext{
			Context:        ctx,
			responseWriter: w,
			request:        r,
			params:         params, // Initialize params
		}

		// Call the handler and handle errors.
		err := handler(c)
		if err != nil {
			app.errorHandler(err)(w, r)
			return
		}
	})
	return nil
}

// Implement HTTP methods.
func (app *App) Get(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodGet, path, handle, middleware...)
}

func (app *App) Post(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodPost, path, handle, middleware...)
}

func (app *App) Options(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodOptions, path, handle, middleware...)
}

func (app *App) Head(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodHead, path, handle, middleware...)
}

func (app *App) Patch(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodPatch, path, handle, middleware...)
}

func (app *App) Delete(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodDelete, path, handle, middleware...)
}

func (app *App) Put(path string, handle HandlerFuncWithContext, middleware ...MiddlewareFuncWithContext) error {
	return app.registerRoute(http.MethodPut, path, handle, middleware...)
}

func (app *App) SetNotFoundHandler(handler HandlerFuncWithContext) {
	app.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(&StdContext{Context: r.Context(), responseWriter: w, request: r})
	})
}

func (app *App) SetNotAllowedHandler(handler HandlerFuncWithContext) {
	app.router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(&StdContext{Context: r.Context(), responseWriter: w, request: r})
	})
}
