package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID      uint
	Name    string
	Players []*Player `gorm:"many2many:player_team_histories;"`
}

// err := db.SetupJoinTable(&Team{}, "Players", &PlayerTeamHistory{})
