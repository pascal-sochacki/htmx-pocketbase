package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pascal-sochacki/htmx-pocketbase.git/views"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.GET("/", func(c echo.Context) error {
			name := c.QueryParam("name")
			component := views.Indexindex(name)
			component.Render(context.Background(), c.Response().Writer)

			return nil

		})

		e.Router.GET("/api/autocomplete", func(c echo.Context) error {

			return c.HTML(http.StatusOK, "<option value=\"test\">")

		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
