package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PlayerTeamHistory struct {
	gorm.Model
	PlayerID  int
	TeamID    int
	StartedAt datatypes.Date
	EndedAt   datatypes.Date
}
