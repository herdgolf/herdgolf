package services

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name      string `json:"name"`
	ProfileID int
	Profile   Profile
}

type Profile struct {
	Link string `json:"link"`
	ID   int
}

func NewServicesPlayer(p Player, db *gorm.DB) *ServicesPlayer {
	return &ServicesPlayer{
		Player: p,
		DB:     db,
	}
}

type ServicesPlayer struct {
	Player Player
	DB     *gorm.DB
}

func (sp *ServicesPlayer) GetAllPlayers() ([]*Player, error) {
	var players []*Player

	if res := sp.DB.Find(&players); res.Error != nil {
		return nil, res.Error
	}

	return players, nil
}

func (sp *ServicesPlayer) GetPlayerById(id int) (Player, error) {
	var players []*Player

	if res := sp.DB.Find(&players, id); res.Error != nil {
		return Player{}, res.Error
	}

	return *players[0], nil
}

func ConverDateTime(tz string, dt time.Time) string {
	loc, _ := time.LoadLocation(tz)

	return dt.In(loc).Format(time.RFC822Z)
}
