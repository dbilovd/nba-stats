package models

import "gorm.io/gorm"

type PlayerGameStat struct {
	gorm.Model    `json:"-"`
	GameID        int     `json:"game_id"`
	PlayerID      int     `json:"player_id"`
	Points        int     `json:"points"`
	Rebounds      int     `json:"rebounds"`
	Assists       int     `json:"assists"`
	Steals        int     `json:"steals"`
	Blocks        int     `json:"blocks"`
	Turnovers     int     `json:"turnovers"`
	Fouls         int     `json:"fouls"`
	MinutesPlayed float64 `json:"minutes_played" gorm:"type:number(5,2)"`
	Player        Player
	Game          Game
}
