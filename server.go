package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"strconv"
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

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {

	fmt.Println("server is running")
	hashval := strconv.FormatUint(uint64(hash("ehsan")), 10)
	fmt.Printf("%T", hashval)
	_, err := connection.Connect()
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
		connection.Connect()
		fmt.Println("Created database")
	}
}
