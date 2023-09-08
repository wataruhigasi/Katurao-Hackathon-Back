package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/handler"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/usecase"
)

func main() {
	conn := &sql.DB{}

	ar := infra.NewArticleRepository(conn)
	au := usecase.NewArticleUsecase(ar)
	ah := handler.NewArticleHandler(au)

	e := echo.New()

	e.GET("/", hello)
	e.GET("/articles", ah.GetArticles())
	e.POST("/article", handler.CreateArticle)

	log.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
