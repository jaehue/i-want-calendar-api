package main

import (
	"log"
	"os"
	"time"

	"github.com/jaehue/i-want-calendar-api/controllers"
	"github.com/jaehue/i-want-calendar-api/factory"
	"github.com/jaehue/i-want-calendar-api/models"
	"github.com/pangpanglabs/goutils/jwtutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtutil.SetJwtSecret(jwtSecret)
	jwtutil.SetExpDuration(time.Hour * 24 * 7)

	db := initDB()

	e := echo.New()

	controllers.HomeController{}.Init(e.Group("/v1"))
	controllers.MemberController{}.Init(e.Group("/v1/members"))

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(factory.ContextDB(db))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwtSecret),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/v1/ping" {
				return true
			}
			if c.Path() == "/v1/login" {
				return true
			}
			return false
		},
	}))

	logrus.SetLevel(logrus.InfoLevel)

	if err := e.Start(":8000"); err != nil {
		log.Println(err)
	}

}
func initDB() *xorm.Engine {
	db, err := xorm.NewEngine("mysql", os.Getenv("DATABASE_CONN"))
	if err != nil {
		panic(err)
	}
	if err := models.Init(db); err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour * 24)
	// db.ShowSQL()

	return db
}
