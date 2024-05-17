package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	ID    uint
	Name  string
	Teams []*Team `gorm:"many2many:player_team_histories;"`
}

// err := db.SetupJoinTable(&Player{}, "Teams", &PlayerTeamHistory{})
