package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type CommentHandler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type commentHandlerImpl struct {
	cr infra.CommentRepository
}

func NewCommentHandler(cr infra.CommentRepository) *commentHandlerImpl {
	return &commentHandlerImpl{
		cr: cr,
	}
}

func (ch *commentHandlerImpl) GetAll(c echo.Context) error {
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

func (ch *commentHandlerImpl) Create(c echo.Context) error {
	id := c.Param("thread_id")
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	idInt, err := strconv.Atoi(id)
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
