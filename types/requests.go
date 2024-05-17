package types

import "encoding/json"

type TeamRequest struct {
	Name string `json:"name"`
}

type PlayerGameStatRequest struct {
	GameID        json.Number `json:"game_id" validate:"required"`
	PlayerID      json.Number `json:"player_id" validate:"required"`
	Points        json.Number `json:"points" validate:"required,gte=0"`
	Rebounds      json.Number `json:"rebounds" validate:"required,gte=0"`
	Assists       json.Number `json:"assists" validate:"required,gte=0"`
	Steals        json.Number `json:"steals" validate:"required,gte=0"`
	Blocks        json.Number `json:"blocks" validate:"required,gte=0"`
	Turnovers     json.Number `json:"turnovers" validate:"required,gte=0"`
	Fouls         json.Number `json:"fouls" validate:"required,gte=0,lte=6"`
	MinutesPlayed json.Number `json:"minutes_played" validate:"required,gte=0,lte=48"`
}
