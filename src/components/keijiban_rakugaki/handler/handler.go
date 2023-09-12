package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/keijiban_rakugaki/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type Handler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type handlerImpl struct {
	r infra.Repo
}

func New(r infra.Repo) *handlerImpl {
	return &handlerImpl{
		r: r,
	}
}

func (h *handlerImpl) GetAll(c echo.Context) error {
	keijiban_rakugakis, err := h.r.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, keijiban_rakugakis)
}

func (h *handlerImpl) Create(c echo.Context) error {
	kr := new(model.KeijibanRakugaki)
	if err := c.Bind(kr); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.r.Create(kr); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
