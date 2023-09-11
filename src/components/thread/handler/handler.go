package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wataruhigasi/Katurao-Hackathon-Back/components/thread/usecase"

	"github.com/wataruhigasi/Katurao-Hackathon-Back/domain/model"
)

type Handler interface {
	GetAll() echo.HandlerFunc
	Create(c echo.Context) error
}

type handlerImpl struct {
	tu usecase.Usecase
}

func New(tu usecase.Usecase) *handlerImpl {
	return &handlerImpl{
		tu: tu,
	}
}

func (th *handlerImpl) GetAll(c echo.Context) error {
	threads, err := th.tu.FindAll()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, threads)
}

type createReq struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func (th *handlerImpl) Create(c echo.Context) error {
	req := new(createReq)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	t := &model.Thread{
		Title: req.Title,
	}

	com := &model.Comment{
		Body:   req.Body,
		Author: req.Author,
	}

	if err := th.tu.Create(t, com); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
