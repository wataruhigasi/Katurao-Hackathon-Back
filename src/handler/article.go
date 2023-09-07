package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/usecase"
)

var _ echo.HandlerFunc = GetArticles

func GetArticles(c echo.Context) error {
	articles, err := usecase.GetArticles(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, articles)
}

var _ echo.HandlerFunc = CreateArticle

func CreateArticle(c echo.Context) error {
	if err := usecase.CreateArticle(c); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
