package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/article/infra"
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
	articles, err := h.r.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, articles)
}

func (h *handlerImpl) Create(c echo.Context) error {
	a := new(model.Article)
	if err := c.Bind(a); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.r.Create(a); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *handlerImpl) ChangePosition(c echo.Context) error {
	id := c.Param("article_id")

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	p := new(model.Position)
	if err := c.Bind(p); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.r.ChangePosition(idInt, p); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
