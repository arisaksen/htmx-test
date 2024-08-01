package main

import (
	"github.com/arisaksen/htmx-test/author"
	"github.com/arisaksen/htmx-test/renderer"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

var (
	runtimeEnvironment string
	templatePath       string
)

func getAuthors() ([]author.Author, error) {
	authors := []author.Author{
		{Name: "J.R.R Tolkien", YearOfBirth: 1982},
		{Name: "Dan Brown", YearOfBirth: 1964},
		{Name: "Sun Tzu", YearOfBirth: -544},
	}
	return authors, nil
}

func handleGetAuthorHtmx(c echo.Context) error {
	authorResponse, err := getAuthors()
	if err != nil {
		return err
	}
	firstAuthor := authorResponse[0]
	response := c.Render(http.StatusOK, "index", firstAuthor)
	return response
}

func init() {
	runtimeEnvironment = os.Getenv("ENVIRONMENT")
	if runtimeEnvironment == "" {
		runtimeEnvironment = "LOCALHOST"
	}

	// HTMX templates
	if runtimeEnvironment == "DOCKER" {
		templatePath = "go/bin/public/*.html"
	} else {
		templatePath = "public/*.html"
	}
}

func main() {
	e := echo.New()
	renderer.NewTemplateRenderer(e, templatePath)
	e.GET("/", handleGetAuthorHtmx)
	e.Logger.Fatal(e.Start(":8080"))
}
