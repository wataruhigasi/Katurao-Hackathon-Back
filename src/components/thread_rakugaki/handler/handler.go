package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread_rakugaki/infra"
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
	id := c.Param("thread_id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	thread_rakugakis, err := h.r.FindAll(idInt)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, thread_rakugakis)
}

func (h *handlerImpl) Create(c echo.Context) error {
	id := c.Param("thread_id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	tr := new(model.ThreadRakugaki)
	if err := c.Bind(tr); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.r.Create(tr, idInt); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
