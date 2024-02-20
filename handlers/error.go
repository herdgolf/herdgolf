package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/herdgolf/herdgolf/views/error_pages"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	var errorPage func(fp bool) templ.Component

	switch code {
	case 401:
		errorPage = error_pages.Error401
	case 404:
		errorPage = error_pages.Error404
	case 500:
		errorPage = error_pages.Error500
	}

	isError = true

	renderView(c, error_pages.ErrorIndex(
		fmt.Sprintf("| Error (%d)", code),
		"",
		fromProtected,
		isError,
		errorPage(fromProtected),
	))
	// errorPage := fmt.Sprintf("views/%d.html", code)
	// if err := c.File(errorPage); err != nil {
	// c.Logger().Error(err)
	// }
}
