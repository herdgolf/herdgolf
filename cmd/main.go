package main

import (
	"net/http"

	"github.com/herdgolf/herdgolf/db"
	"github.com/herdgolf/herdgolf/handlers"
	"github.com/herdgolf/herdgolf/services"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	app.Static("/", "assets")

	app.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/player")
	})

	db.Init()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	dbGorm.Ping()

	ps := services.NewServicesPlayer(services.Player{}, gorm)

	p := handlers.New(ps)

	handlers.SetupRoutes(app, p)

	app.Logger.Fatal(app.Start(":8080"))
}
