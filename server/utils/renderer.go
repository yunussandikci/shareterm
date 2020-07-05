package utils

import (
	"fmt"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewRenderer(path string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(fmt.Sprintf("%s/*.html", path))),
	}
}
