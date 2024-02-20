package main

import (
	"github.com/gorilla/sessions"
	"github.com/herdgolf/herdgolf/db"
	"github.com/herdgolf/herdgolf/handlers"
	"github.com/herdgolf/herdgolf/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SECRET_KEY = "secret"
)

func main() {
	e := echo.New()

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	e.Static("/", "assets")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

	db.Init()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	dbGorm.Ping()

	ps := services.NewServicesPlayer(services.Player{}, gorm)
	ah := handlers.NewAuthHandler(ps)

	p := handlers.New(ps)

	handlers.SetupRoutes(e, p, ah)

	e.Logger.Fatal(e.Start(":8080"))
}
