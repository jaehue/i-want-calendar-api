package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/jwtutil"
)

type HomeController struct{}

func (c HomeController) Init(g *echo.Group) {
	g.GET("/ping", c.Ping)
	g.POST("/login", c.Login)
}
func (this HomeController) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
func (this HomeController) Login(c echo.Context) error {
	var body struct {
		Password string
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, ApiResult{
			Error: ApiError{
				Message: err.Error(),
			},
		})
	}
	password := os.Getenv("LOGiN_PASSWORD")
	if body.Password != password {
		return c.JSON(http.StatusBadRequest, ApiResult{
			Error: ApiError{
				Message: "Invalid password",
			},
		})
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
			"token": token,
		},
	})
}
