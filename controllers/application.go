package controllers

import (
	"net/http"
	"strconv"

	"github.com/jaehue/i-want-calendar-api/models"
	"github.com/labstack/echo"
)

type ApplicationController struct{}

func (c ApplicationController) Init(g *echo.Group) {
	g.GET("", c.GetApplications)
	g.POST("", c.CreateApplication)
	g.PUT("/:id", c.UpdateApplication)
}

func (this ApplicationController) GetApplications(c echo.Context) error {
	memberId, _ := strconv.ParseInt(c.QueryParam("memberId"), 10, 64)
	applications, err := models.Application{}.GetAll(c.Request().Context(), memberId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, ApiResult{
		Success: true,
		Result:  applications,
	})
}

func (this ApplicationController) CreateApplication(c echo.Context) error {
	var application models.Application
	if err := c.Bind(&application); err != nil {
		return c.JSON(http.StatusBadRequest, ApiResult{Error: ApiError{Message: err.Error()}})
	}

	if err := application.Create(c.Request().Context()); err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, ApiResult{
		Success: true,
		Result:  application,
	})
}

func (this ApplicationController) UpdateApplication(c echo.Context) error {
	applicationId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if applicationId == 0 {
		return c.JSON(http.StatusBadRequest, ApiResult{Error: ApiError{Message: "Invalid application id"}})
	}
	var application models.Application
	if err := c.Bind(&application); err != nil {
		return c.JSON(http.StatusBadRequest, ApiResult{Error: ApiError{Message: err.Error()}})
	}
	application.Id = applicationId

	if err := application.Update(c.Request().Context()); err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, ApiResult{
		Success: true,
		Result:  application,
	})
}
