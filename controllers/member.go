package controllers

import (
	"net/http"
	"strconv"

	"github.com/jaehue/i-want-calendar-api/models"

	"github.com/labstack/echo"
)

type MemberController struct{}

func (c MemberController) Init(g *echo.Group) {
	g.GET("", c.GetMembers)
	g.POST("", c.CreateMember)
}

func (this MemberController) GetMembers(c echo.Context) error {
	memberType, _ := strconv.ParseInt(c.QueryParam("memberType"), 10, 64)
	teacherId, _ := strconv.ParseInt(c.QueryParam("teacherId"), 10, 64)
	includeGraduate, _ := strconv.ParseBool(c.QueryParam("includeGraduate"))

	members, err := models.Member{}.GetAll(c.Request().Context(), models.MemberType(memberType), teacherId, includeGraduate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, ApiResult{
		Success: true,
		Result:  members,
	})
}

func (this MemberController) CreateMember(c echo.Context) error {
	var member models.Member
	if err := c.Bind(&member); err != nil {
		return c.JSON(http.StatusBadRequest, ApiResult{Error: ApiError{Message: err.Error()}})
	}

	if err := member.Create(c.Request().Context()); err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}
	return c.JSON(http.StatusOK, ApiResult{
		Success: true,
		Result:  member,
	})
}
