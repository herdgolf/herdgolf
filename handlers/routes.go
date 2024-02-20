package handlers

import "github.com/labstack/echo/v4"

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(e *echo.Echo, p *PlayerHandler, ah *AuthHandler) {
	e.GET("/", ah.homeHandler)
	e.GET("/login", ah.loginHandler)
	e.POST("/login", ah.loginHandler)
	e.GET("/register", ah.registerHandler)
	e.POST("/register", ah.registerHandler)
	protectedGroup := e.Group("/player", ah.authMiddleware)
	protectedGroup.GET("", p.HandlerShowPlayers)
	protectedGroup.GET("/details/:id", p.HandlerShowPlayerById)
}
