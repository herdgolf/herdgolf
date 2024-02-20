package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/herdgolf/herdgolf/services"
	"github.com/herdgolf/herdgolf/views/player"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PlayerService interface {
	GetAllPlayers() ([]*services.Player, error)
	GetPlayerById(id int) (services.Player, error)
}

func New(ps PlayerService) *PlayerHandler {
	return &PlayerHandler{ps}
}

type PlayerHandler struct {
	PlayerService PlayerService
}

func (ph *PlayerHandler) HandlerShowPlayers(c echo.Context) error {
	players, err := ph.PlayerService.GetAllPlayers()
	if err != nil {
		return err
	}

	si := player.ShowIndex("| Home", player.Show(players))
	return ph.View(c, si)
}

func (ph *PlayerHandler) HandlerShowPlayerById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	pdata, err := ph.PlayerService.GetPlayerById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return err
	}

	di := player.DetailsIndex(
		fmt.Sprintf("| User details %s",
			cases.Title(language.English).String(pdata.Name),
		),
		player.Details(tz, pdata),
	)

	return ph.View(c, di)
}

func (ph *PlayerHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
