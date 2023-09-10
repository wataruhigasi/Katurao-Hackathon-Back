package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/infra"

	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ThreadHandler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type threadHandlerImpl struct {
	tr infra.ThreadRepository
}

func NewThreadHandler(tr infra.ThreadRepository) *threadHandlerImpl {
	return &threadHandlerImpl{
		tr: tr,
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

	res, err := th.tr.Create(t)

		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}

// comment := new(model.Comment)
// if err := c.Bind(comment); err != nil {
// 	c.Logger().Error(err)
// 	return echo.NewHTTPError(http.StatusBadRequest, err)
// }

// if err := th.cr.Create(comment, idInt); err != nil {
// 	c.Logger().Error(err)
// 	return echo.NewHTTPError(http.StatusInternalServerError, err)
// }
