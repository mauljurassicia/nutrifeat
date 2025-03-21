package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type App struct {
	router *http.ServeMux
	errorHandler NewErrorHandler
}

func NewStd() *App {
	return &App{router: http.NewServeMux(),
		errorHandler: func(err error) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(err.Error()))
			}
		},
	}
}

type StdContext struct {
	Context *context.Context // Embed the context.Context
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
	fmt.Println(string(body))
	_, err := c.responseWriter.Write(body)
	return err
}

func (c *StdContext) SendStatus(status int) error {
	c.responseWriter.WriteHeader(status)
	return nil
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

func (c *StdContext) BodyParser(body interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(body)
}

func (c *StdContext) Render(template string, data interface{}, layout ...string) error {
	// Rendering logic should be implemented here, depending on the template engine used
	return nil
}



// Serve starts the HTTP server.
func (app *App) Serve(addr string) error {
	return http.ListenAndServe(addr, app.router)
}

func (app *App) stdHandler(h HandlerFuncWithContext) http.HandlerFunc {

	fmt.Println(h)
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Println(ctx)

		// Create a new StdContext.
		c := &StdContext{
			Context:        &ctx,
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

// Register HTTP methods dynamically.
func (app *App) registerRoute(method string, path string, handle HandlerFuncWithContext) error {
	app.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// ✅ Call stdHandler(handle) and execute its result
		app.stdHandler(handle)(w, r)
	})
	return nil
}


// Implement HTTP methods.
func (app *App) Get(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodGet, path, handle)
}

func (app *App) Post(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodPost, path, handle)
}

func (app *App) Options(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodOptions, path, handle)
}

func (app *App) Head(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodHead, path, handle)
}

func (app *App) Patch(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodPatch, path, handle)
}

func (app *App) Delete(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodDelete, path, handle)
}

func (app *App) Put(path string, handle HandlerFuncWithContext) error {
	return app.registerRoute(http.MethodPut, path, handle)
}