package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type Handler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type handlerImpl struct {
	ar infra.Repo
}

func New(ar infra.Repo) *handlerImpl {
	return &handlerImpl{
		ar: ar,
	}
}

func (ah *handlerImpl) GetAll(c echo.Context) error {
	articles, err := ah.ar.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (ah *handlerImpl) Create(c echo.Context) error {
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
