package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ArticleHandler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type articleHandlerImpl struct {
	ar infra.ArticleRepository
}

func NewArticleHandler(ar infra.ArticleRepository) *articleHandlerImpl {
	return &articleHandlerImpl{
		ar: ar,
	}
}

func (ah *articleHandlerImpl) GetAll(c echo.Context) error {
	articles, err := ah.ar.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (ah *articleHandlerImpl) Create(c echo.Context) error {
	a := new(model.Article)
	if err := c.Bind(a); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ah.ar.Create(a); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
