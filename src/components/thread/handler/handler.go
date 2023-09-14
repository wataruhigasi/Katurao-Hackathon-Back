package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type Handler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
	ChangePosition(c echo.Context) error
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
	threads, err := h.r.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, threads)
}

func (h *handlerImpl) Create(c echo.Context) error {
	t := new(model.Thread)
	if err := c.Bind(t); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.r.Create(t); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *handlerImpl) ChangePosition(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("thread_id"), 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := new(model.Position)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.r.ChangePosition(id, req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
