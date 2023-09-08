package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/usecase"
)

type ArticleHandler interface {
	GetArticles() echo.HandlerFunc
}

type articleHandlerImpl struct {
	au usecase.ArticleUsecase
}

func NewArticleHandler(au usecase.ArticleUsecase) *articleHandlerImpl {
	return &articleHandlerImpl{
		au: au,
	}
}

func (ah *articleHandlerImpl) GetArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		articles, err := ah.au.GetArticles()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, articles)
	}
}

var _ echo.HandlerFunc = CreateArticle

func CreateArticle(c echo.Context) error {
	if err := usecase.CreateArticle(); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
