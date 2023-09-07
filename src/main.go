package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/handler"
)

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/articles", handler.GetArticles)
	e.POST("/article", handler.CreateArticle)

	log.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
