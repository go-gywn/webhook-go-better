package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"time"
	"webhook-better/config"
	"webhook-better/middlewares"
)

func main() {
	config.InitConfig("")

	e := echo.New()
	e.Use(middleware.CORS())
	e.HideBanner = true

	dsn := sqlite.Open("account.db")
	db, err := gorm.Open(dsn, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Database connection error")
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	e.Use(middlewares.JwtToken())
	e.Use(middlewares.GORM(db))

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})
	e.POST("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})
	e.Start(config.Config.Bind)
}
