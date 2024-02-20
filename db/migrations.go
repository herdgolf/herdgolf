package db

import (
	"github.com/herdgolf/herdgolf/services"
)

func AutoMigrate() {
	// database.AutoMigrate(&models.Profile{})
	database.AutoMigrate(&services.Player{})
	// database.AutoMigrate(&models.Hole{})
	// database.AutoMigrate(&models.Round{})
	// database.AutoMigrate(&models.ScoreCard{})
	// database.AutoMigrate(&models.HoleScore{})
	// database.AutoMigrate(&models.Course{})
}
