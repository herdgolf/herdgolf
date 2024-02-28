package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/herdgolf/herdgolf/services"
	"github.com/herdgolf/herdgolf/views/scorecard_views"
	"github.com/labstack/echo/v4"
)

type ScoreCardService interface {
	CreateScoreCard(scoreCard services.ScoreCard) error
	APICreateScoreCard(scoreCard *services.ScoreCard) (*services.ScoreCard, error)
	GetAllScoreCards() ([]*services.ScoreCard, error)
	GetScoreCardById(id int) (services.ScoreCard, error)
	// UpdateCourse(course services.Course) error
}

func NewScoreCardServiceHandler(sc ScoreCardService) *ScoreCardServiceHandler {
	return &ScoreCardServiceHandler{sc}
}

type ScoreCardServiceHandler struct {
	ServiceScoreCard ScoreCardService
}

func (sc *ScoreCardServiceHandler) CreateScoreCard(c echo.Context) error {
	isError = false

	if c.Request().Method == "POST" {
		p, _ := strconv.Atoi(c.FormValue("par"))
		slope, _ := strconv.Atoi(c.FormValue("slope"))
		courseId, _ := strconv.ParseUint(c.FormValue("courseId"), 10, 64)
		rating, _ := strconv.ParseFloat(c.FormValue("rating"), 64)
		scoreCard := services.ScoreCard{
			Tee:      c.FormValue("tee"),
			Par:      p,
			Rating:   rating,
			Slope:    slope,
			CourseID: uint(courseId),
		}

		fmt.Printf("Handler scorecard: %v\n", scoreCard)

		err := sc.ServiceScoreCard.CreateScoreCard(scoreCard)
		if err != nil {
			return err
		}

		setFlashmessages(c, "success", "Course created successfully!")

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/course/details/%d", 1))
	}

	return renderView(c, scorecard_views.ScoreCardIndex(
		"| Create Score Card",
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		scorecard_views.CreateScoreCard(),
	))
}

func (sc *ScoreCardServiceHandler) listScoreCards(c echo.Context) error {
	isError = false

	scoreCards, err := sc.ServiceScoreCard.GetAllScoreCards()
	if err != nil {
		return err
	}

	return renderView(c, scorecard_views.ScoreCardIndex(
		"| Score Card List",
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		scorecard_views.ScoreCardList(scoreCards),
	))
}

func (sc *ScoreCardServiceHandler) HandlerShowScoreCardById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	// cdata, err := ph.PlayerService.GetPlayerById(idParam)
	cdata, err := sc.ServiceScoreCard.GetScoreCardById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return err
	}

	return renderView(c, scorecard_views.Details(tz, cdata))

	// return renderView(c, scorecard_views.DetailsIndex(
	// 	fmt.Sprintf("| Scorecard details %s",
	// 		cases.Title(language.English).String(cdata.Tee),
	// 	),
	// 	"",
	// 	fromProtected,
	// 	isError,
	// 	getFlashmessages(c, "error"),
	// 	getFlashmessages(c, "success"),
	// 	scorecard_views.Details(tz, cdata),
	// ))
}

func (sc *ScoreCardServiceHandler) APICreateScoreCard(c echo.Context) error {
	cs := new(services.ScoreCard)

	if err := c.Bind(cs); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	scorecard := &services.ScoreCard{
		Tee:      cs.Tee,
		Par:      cs.Par,
		Rating:   cs.Rating,
		Slope:    cs.Slope,
		CourseID: cs.CourseID,
		Holes:    cs.Holes,
	}
	card, err := sc.ServiceScoreCard.APICreateScoreCard(scorecard)
	if err != nil {
		return err
	}

	response := map[string]interface{}{
		"data": &card,
	}

	return c.JSON(http.StatusCreated, response)
}

// func (sc *ScoreCardServiceHandler) HandlerShowCourseById(c echo.Context) error {
// 	idParam, _ := strconv.Atoi(c.Param("id"))
//
// 	tz := ""
// 	if len(c.Request().Header["X-Timezone"]) != 0 {
// 		tz = c.Request().Header["X-Timezone"][0]
// 	}
//
// 	// cdata, err := ph.PlayerService.GetPlayerById(idParam)
// 	cdata, err := cs.ServiceCourse.GetCourseById(idParam)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "no rows in result set") {
// 			return echo.NewHTTPError(http.StatusNotFound, err)
// 		}
// 		return err
// 	}
//
// 	return renderView(c, course_views.DetailsIndex(
// 		fmt.Sprintf("| Course details %s",
// 			cases.Title(language.English).String(cdata.Name),
// 		),
// 		"",
// 		fromProtected,
// 		isError,
// 		getFlashmessages(c, "error"),
// 		getFlashmessages(c, "success"),
// 		course_views.Details(tz, cdata),
// 	))
// }
