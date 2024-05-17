package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID         uint
	SeasonID   int
	Season     Season
	HomeTeamID int
	HomeTeam   Team
	AwayTeamID int
	AwayTeam   Team
	Date       datatypes.Date
}
