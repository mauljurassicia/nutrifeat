package config

import (
	"github/mauljurassicia/nutrifeat/infrastructure/http"
	"github/mauljurassicia/nutrifeat/model"
	pages "github/mauljurassicia/nutrifeat/presentation/view/pages/not_found"

	"github.com/spf13/viper"
)

func NewServer(config *viper.Viper) http.HttpApp {

	app := http.NewStd()

	app.SetNotFoundHandler(func(c http.HttpContext) error {
		if c.Header("Content-Type") == "application/json" {
			return c.Status(http.NotFound).JSON(model.NewErrorResponse(c, http.NotFound, model.ErrNotFound, "Not Found"))
		}

		return c.Render(pages.NotFound())
	})

	return app
}
