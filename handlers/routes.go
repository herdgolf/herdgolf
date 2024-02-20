package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo, p *PlayerHandler) {
	group := app.Group("/player")
	group.GET("", p.HandlerShowPlayers)
	group.GET("/details/:id", p.HandlerShowPlayerById)
}
