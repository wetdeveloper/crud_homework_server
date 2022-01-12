package main

import (
	"fmt"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/wetdeveloper/connection"
	"github.com/wetdeveloper/crud_api"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	fmt.Println("server is running")
	_, err := connection.Connect()
	// connection.CreateTable(mydb)
	// connection.InsertUser(mydb, "ehsan", "Heydary")
	if err == nil {
		fmt.Println("You are connected to database")
		e := echo.New()
		renderer := &TemplateRenderer{
			templates: template.Must(template.ParseGlob("*.html")),
		}
		e.Renderer = renderer

		e.GET("/userslist/", crud_api.Read)
		e.GET("", crud_api.CrudForm)
		e.POST("/crud-page/crud/", crud_api.Cud)
		e.Start(":")
	} else {
		connection.CreateDb()
		fmt.Println("Created database")
	}
}
