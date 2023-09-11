package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type Handler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type handlerImpl struct {
	cr infra.Repo
}

func New(cr infra.Repo) *handlerImpl {
	return &handlerImpl{
		cr: cr,
	}
}

func (ch *handlerImpl) GetAll(c echo.Context) error {
	id := c.Param("thread_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	comments, err := ch.cr.FindAll(idInt)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, comments)
}

func (ch *handlerImpl) Create(c echo.Context) error {
	id := c.Param("thread_id")
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ch.cr.Create(comment, idInt); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
