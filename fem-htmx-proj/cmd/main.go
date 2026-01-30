package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Templates wraps Go's template engine for Echo
type Templates struct {
	templates *template.Template
}

// Count holds the counter state
type Count struct {
	Count int
}

// Render implements Echo's Renderer interface to render HTML templates
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplates loads all HTML templates from the views directory
func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()             // Create Echo web server instance
	e.Use(middleware.Logger())  // Add logging middleware
	count := Count{Count: 0}    // Initialize counter
	e.Renderer = NewTemplates() // Register template renderer

	// Handle GET requests to root path
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++ // Increment counter on each request render only the count blockj
		return c.Render(http.StatusOK, "count", count)
	})
	e.Logger.Fatal(e.Start(":42069")) // Start server on port 42069
}
