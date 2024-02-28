package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/herdgolf/herdgolf/services"
	"github.com/herdgolf/herdgolf/views/course_views"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CourseService interface {
	CreateCourse(course services.Course) error
	GetAllCourses() ([]*services.Course, error)
	GetCourseById(id int) (services.Course, error)
	UpdateCourse(course services.Course) error
}

func NewCourseServiceHandler(cs CourseService) *CourseServiceHandler {
	return &CourseServiceHandler{cs}
}

type CourseServiceHandler struct {
	ServiceCourse CourseService
}

func (cs *CourseServiceHandler) createCourse(c echo.Context) error {
	isError = false

	if c.Request().Method == "POST" {
		p, _ := strconv.Atoi(c.FormValue("par"))
		course := services.Course{
			Name: strings.Trim(c.FormValue("name"), " "),
			Par:  p,
		}

		err := cs.ServiceCourse.CreateCourse(course)
		if err != nil {
			return err
		}

		setFlashmessages(c, "success", "Course created successfully!")

		return c.Redirect(http.StatusSeeOther, "/course")
	}

	return renderView(c, course_views.CourseIndex(
		"| Create Course",
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		course_views.CreateCourse(),
	))
}

func (cs *CourseServiceHandler) listCourses(c echo.Context) error {
	isError = false
	courses, err := cs.ServiceCourse.GetAllCourses()
	if err != nil {
		return err
	}

	return renderView(c, course_views.CourseIndex(
		"| Course List",
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		course_views.CourseList(courses),
	))
}

func (cs *CourseServiceHandler) HandlerShowCourseById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	// cdata, err := ph.PlayerService.GetPlayerById(idParam)
	cdata, err := cs.ServiceCourse.GetCourseById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return err
	}

	// di := player.DetailsIndex(
	// 	fmt.Sprintf("| User details %s",
	// 		cases.Title(language.English).String(pdata.Name),
	// 	),
	// 	player.Details(tz, pdata),
	// )

	// return ph.View(c, di)
	return renderView(c, course_views.DetailsIndex(
		fmt.Sprintf("| Course details %s",
			cases.Title(language.English).String(cdata.Name),
		),
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		course_views.Details(tz, cdata),
	))
}

func (cs *CourseServiceHandler) updateCourse(c echo.Context) error {
	isError = false

	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	fmt.Println(idParam)

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}
	cdata, err := cs.ServiceCourse.GetCourseById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}
	fmt.Println(cdata)

	if c.Request().Method == "POST" {
		p, _ := strconv.Atoi(c.FormValue("par"))
		cdata.Name = c.FormValue("name")
		cdata.Par = p

		fmt.Println(cdata)
		err := cs.ServiceCourse.UpdateCourse(cdata)
		if err != nil {
			return err
		}

		setFlashmessages(c, "success", "Task created successfully!!")

		return c.Redirect(http.StatusSeeOther, "/course/details/"+strconv.Itoa(idParam))
	}

	return renderView(c, course_views.CourseIndex(
		fmt.Sprintf("Edit Course %s", cdata.Name),
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		course_views.UpdateCourse(cdata, tz),
		// course_views.UpdateCourse(cdata, c.Get(tzone_key).(string)),
	))
}
