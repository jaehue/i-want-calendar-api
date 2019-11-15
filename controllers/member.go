package controllers

import (
	"net/http"
	"strconv"

	"github.com/jaehue/i-want-calendar-api/models"
	"github.com/pangpanglabs/goutils/jwtutil"

	"github.com/labstack/echo"
)

type MemberController struct{}

func (c MemberController) Init(g *echo.Group) {
	g.GET("", c.GetMembers)
	g.POST("/login/facebook", c.LoginFacebookMember)
	g.POST("", c.CreateMember)
}
func (this MemberController) LoginFacebookMember(c echo.Context) error {
	var p struct {
		AccessToken              string `json:"accessToken"`
		DataAccessExpirationTime int64  `json:"data_access_expiration_time"`
		Email                    string `json:"email"`
		ExpiresIn                int64  `json:"expiresIn"`
		Name                     string `json:"name"`
		SignedRequest            string `json:"signedRequest"`
		UserID                   string `json:"userID"`
	}
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}

	if p.UserID == "" {
		return c.JSON(http.StatusBadRequest, ApiResult{
			Error: ApiError{
				Message: "Invalid facebook user information",
			},
		})
	}

	member, err := models.Member{}.GetByFacebookUserId(c.Request().Context(), p.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}

	if member == nil {
		member = &models.Member{
			Name:              p.Name,
			Email:             p.Email,
			FacebookUserId:    p.UserID,
			FacebookExpiresIn: p.ExpiresIn,
		}
		if err := member.Create(c.Request().Context()); err != nil {
			return c.JSON(http.StatusInternalServerError, ApiResult{
				Error: ApiError{
					Message: err.Error(),
				},
			})
		}
	}

	token, err := jwtutil.NewToken(nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}

	return c.JSON(http.StatusOK, ApiResult{
		Success: true,
		Result: map[string]interface{}{
			"token":  token,
			"member": member,
		},
	})

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
