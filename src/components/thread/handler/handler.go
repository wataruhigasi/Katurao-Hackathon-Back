package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	infrac "github.com/wataruhigasi/Katurao-Hackathon-Back/components/comment/infra"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"

	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ThreadHandler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type threadHandlerImpl struct {
	tr infra.ThreadRepository
	cr infrac.CommentRepository
}

func NewThreadHandler(tr infra.ThreadRepository, cr infrac.CommentRepository) *threadHandlerImpl {
	return &threadHandlerImpl{
		tr: tr,
		cr: cr,
	}
}

func (th *threadHandlerImpl) GetAll(c echo.Context) error {
	threads, err := th.tr.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, threads)
}

func (th *threadHandlerImpl) Create(c echo.Context) error {
	t := new(model.Thread)

	if err := c.Bind(t); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// ここから

	if err := th.tr.Create(t); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	id, err := th.tr.GetLastInsertID()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// ここまで

	// id を int に変換
	idInt, err := strconv.Atoi(strconv.FormatInt(id, 10))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := th.cr.Create(comment, idInt); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
