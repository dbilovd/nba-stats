package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model
	Name      string
	StartedAt datatypes.Date
	EndedAt   datatypes.Date
}
