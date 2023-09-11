package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/usecase"

	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type ThreadHandler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type threadHandlerImpl struct {
	tu usecase.ThreadUsecase
}

func NewThreadHandler(tu usecase.ThreadUsecase) *threadHandlerImpl {
	return &threadHandlerImpl{
		tu: tu,
	}
}

func (th *threadHandlerImpl) GetAll(c echo.Context) error {
	threads, err := th.tu.FindAll()
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

	com := new(model.Comment)
	if err := c.Bind(com); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := th.tu.Create(t, com); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
