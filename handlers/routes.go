package handlers

import "github.com/labstack/echo/v4"

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(
	e *echo.Echo,
	p *PlayerHandler,
	c *CourseServiceHandler,
	s *ScoreCardServiceHandler,
	ah *AuthHandler,
) {
	e.GET("/", ah.homeHandler)
	e.GET("/login", ah.loginHandler)
	e.POST("/login", ah.loginHandler)
	e.GET("/register", ah.registerHandler)
	e.POST("/register", ah.registerHandler)
	apiGroup := e.Group("/api")
	apiGroup.POST("/scorecard", s.APICreateScoreCard)
	protectedGroup := e.Group("/player", ah.authMiddleware)
	protectedGroup.GET("", p.HandlerShowPlayers)
	protectedGroup.GET("/details/:id", p.HandlerShowPlayerById)
	protectedGroup.GET("/edit/:id", p.updatePlayerHandler)
	protectedGroup.POST("/edit/:id", p.updatePlayerHandler)
	protectedGroup.POST("/logout", p.logoutHandler)
	coursesGroup := e.Group("/course", ah.authMiddleware)
	coursesGroup.GET("", c.listCourses)
	coursesGroup.GET("/create", c.createCourse)
	coursesGroup.POST("/create", c.createCourse)
	coursesGroup.GET("/details/:id", c.HandlerShowCourseById)
	coursesGroup.GET("/edit/:id", c.updateCourse)
	coursesGroup.POST("/edit/:id", c.updateCourse)
	scoreCardGroup := e.Group("/scorecard", ah.authMiddleware)
	scoreCardGroup.GET("", s.listScoreCards)
	scoreCardGroup.GET("/create", s.CreateScoreCard)
	scoreCardGroup.POST("/create", s.CreateScoreCard)
	scoreCardGroup.GET("/details/:id", s.HandlerShowScoreCardById)
}
