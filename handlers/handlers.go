package handlers

import (
	"fmt"
	"strconv"
	"takehome/database"
	"takehome/models"
	"takehome/types"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

var validate *validator.Validate

func GetReports(c *fiber.Ctx) error {
	var stats []struct {
		PlayerID      float64 `json:"player_id"`
		Games         float64 `json:"games"`
		Points        float64 `json:"points"`
		Rebounds      float64 `json:"rebounds"`
		Assists       float64 `json:"assists"`
		Steals        float64 `json:"steals"`
		Blocks        float64 `json:"blocks"`
		Turnovers     float64 `json:"turnovers"`
		Fouls         float64 `json:"fouls"`
		MinutesPlayed float64 `json:"minutes_played"`
		PlayerName    string  `json:"player_name"`
	}
	database.DB.Model(&models.PlayerGameStat{}).Select(
		"player_id", "COUNT(game_id) AS games", "Player.name AS player_name", "AVG(points) AS points", "AVG(rebounds) AS rebounds", "AVG(assists) AS assists", "AVG(steals) AS steals", "AVG(blocks) AS blocks", "AVG(turnovers) AS turnovers", "AVG(fouls) AS fouls", "AVG(minutes_played) AS minutes_played",
	).Joins("Player").Group("player_id").Find(&stats)

	var team_stats []struct {
		ID            float64 `json:"id"`
		TeamName      string  `json:"team_name"`
		Games         float64 `json:"games"`
		Points        float64 `json:"points"`
		Rebounds      float64 `json:"rebounds"`
		Assists       float64 `json:"assists"`
		Steals        float64 `json:"steals"`
		Blocks        float64 `json:"blocks"`
		Turnovers     float64 `json:"turnovers"`
		Fouls         float64 `json:"fouls"`
		MinutesPlayed float64 `json:"minutes_played"`
	}
	database.DB.Model(&models.PlayerGameStat{}).Select(
		"teams.id", "teams.name AS team_name", "AVG(points) AS points", "AVG(rebounds) AS rebounds", "AVG(assists) AS assists", "AVG(steals) AS steals", "AVG(blocks) AS blocks", "AVG(turnovers) AS turnovers", "AVG(fouls) AS fouls", "AVG(minutes_played) AS minutes_played",
	).Joins(
		"JOIN player_team_histories ON player_team_histories.player_id = player_game_stats.player_id",
	).Joins(
		"JOIN teams ON teams.id = player_team_histories.team_id",
	).Group("teams.id").Find(&team_stats)

	return c.Render("reports", fiber.Map{
		"Stats":     stats,
		"TeamStats": team_stats,
	})
}

func GetPlayerStats(c *fiber.Ctx) error {
	var stats []struct {
		Player_id     int `json:"player_id"`
		Points        int `json:"points"`
		Rebounds      int `json:"rebounds"`
		Assists       int `json:"assists"`
		Steals        int `json:"steals"`
		Blocks        int `json:"blocks"`
		Turnovers     int `json:"turnovers"`
		Fouls         int `json:"fouls"`
		MinutesPlayed int `json:"minutes_played"`
	}
	database.DB.Model(&models.PlayerGameStat{}).Select(
		"player_id", "AVG(points) AS points", "AVG(rebounds) AS rebounds", "AVG(assists) AS assists", "AVG(steals) AS steals", "AVG(blocks) AS blocks", "AVG(turnovers) AS turnovers", "AVG(fouls) AS fouls", "AVG(minutes_played) AS minutes_played",
	).Group("player_id").Find(&stats)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    stats,
	})
}

func GetTeamStats(c *fiber.Ctx) error {
	var team_stats []struct {
		ID            float64 `json:"id"`
		TeamName      string  `json:"team_name"`
		Games         float64 `json:"games"`
		Points        float64 `json:"points"`
		Rebounds      float64 `json:"rebounds"`
		Assists       float64 `json:"assists"`
		Steals        float64 `json:"steals"`
		Blocks        float64 `json:"blocks"`
		Turnovers     float64 `json:"turnovers"`
		Fouls         float64 `json:"fouls"`
		MinutesPlayed float64 `json:"minutes_played"`
	}
	database.DB.Model(&models.PlayerGameStat{}).Select(
		"teams.id", "teams.name AS team_name", "AVG(points) AS points", "AVG(rebounds) AS rebounds", "AVG(assists) AS assists", "AVG(steals) AS steals", "AVG(blocks) AS blocks", "AVG(turnovers) AS turnovers", "AVG(fouls) AS fouls", "AVG(minutes_played) AS minutes_played",
	).Joins(
		"JOIN player_team_histories ON player_team_histories.player_id = player_game_stats.player_id",
	).Joins(
		"JOIN teams ON teams.id = player_team_histories.team_id",
	).Group("teams.id").Find(&team_stats)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    team_stats,
	})
}

type ValidationError struct {
	Field string
	Error string
}

func PostPlayerStats(c *fiber.Ctx) error {
	request := new(types.PlayerGameStatRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	validate = validator.New()

	err := validate.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		error_messages := make([]ValidationError, len(errors))
		for i, error := range errors {
			error_messages[i] = ValidationError{error.Field(), error.Tag()}
		}

		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": error_messages,
		})
	}

	player_id, _ := request.PlayerID.Int64()
	game_id, _ := request.GameID.Int64()
	points, _ := request.Points.Int64()
	rebounds, _ := request.Rebounds.Int64()
	assists, _ := request.Assists.Int64()
	steals, _ := request.Steals.Int64()
	blocks, _ := request.Blocks.Int64()
	turnovers, _ := request.Turnovers.Int64()
	fouls, _ := request.Fouls.Int64()
	minutes_played, _ := request.MinutesPlayed.Float64()

	game_stat := models.PlayerGameStat{
		PlayerID:      int(player_id),
		GameID:        int(game_id),
		Points:        int(points),
		Rebounds:      int(rebounds),
		Assists:       int(assists),
		Steals:        int(steals),
		Blocks:        int(blocks),
		Turnovers:     int(turnovers),
		Fouls:         int(fouls),
		MinutesPlayed: minutes_played,
	}

	result := database.DB.Create(&game_stat)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    game_stat,
	})
}

func PostTeam(c *fiber.Ctx) error {
	request := new(types.TeamRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	team := models.Team{
		Name: request.Name,
	}

	result := database.DB.Create(&team)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    team,
	})
}

func PostPlayer(c *fiber.Ctx) error {
	type PlayerRequest struct {
		Name string `json:"name"`
	}
	request := new(PlayerRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	player := models.Player{
		Name: request.Name,
	}

	result := database.DB.Create(&player)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    player,
	})
}

func PostPlayerTeamHistory(c *fiber.Ctx) error {
	type PlayerTeamHistoryRequest struct {
		TeamID    int    `json:"team_id"`
		StartedAt string `json:"started_at"`
		EndedAt   string `json:"ended_at"`
	}
	request := new(PlayerTeamHistoryRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	player_id, _ := c.ParamsInt("playerId")
	team_id := request.TeamID
	started_at_parsed, _ := time.Parse(time.DateOnly, request.StartedAt)
	started_at := datatypes.Date(started_at_parsed)
	ended_at_parsed, _ := time.Parse(time.DateOnly, request.EndedAt)
	ended_at := datatypes.Date(ended_at_parsed)

	history := models.PlayerTeamHistory{
		PlayerID:  player_id,
		TeamID:    team_id,
		StartedAt: started_at,
		EndedAt:   ended_at,
	}

	result := database.DB.Create(&history)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    history,
	})
}

func PostSeason(c *fiber.Ctx) error {
	type SeasonRequest struct {
		Name      string `json:"name"`
		StartedAt string `json:"started_at"`
		EndedAt   string `json:"ended_at"`
	}
	request := new(SeasonRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	started_at_parsed, _ := time.Parse(time.DateOnly, request.StartedAt)
	started_at := datatypes.Date(started_at_parsed)
	ended_at_parsed, _ := time.Parse(time.DateOnly, request.EndedAt)
	ended_at := datatypes.Date(ended_at_parsed)
	season := models.Season{
		Name:      request.Name,
		StartedAt: started_at,
		EndedAt:   ended_at,
	}

	result := database.DB.Create(&season)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    season,
	})
}

func PostGames(c *fiber.Ctx) error {
	type GameRequest struct {
		SeasonID   string `json:"season_id"`
		HomeTeamID string `json:"home_team_id"`
		AwayTeamID string `json:"away_team_id"`
		Date       string `json:"date"`
	}

	request := new(GameRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	fmt.Println("up to here!")

	season_id, _ := strconv.Atoi(request.SeasonID)
	home_team_id, _ := strconv.Atoi(request.HomeTeamID)
	away_team_id, _ := strconv.Atoi(request.AwayTeamID)
	parsed_time, _ := time.Parse(time.DateOnly, request.Date)
	match_date := datatypes.Date(parsed_time)
	game := models.Game{
		SeasonID:   season_id,
		HomeTeamID: home_team_id,
		AwayTeamID: away_team_id,
		Date:       match_date,
	}

	result := database.DB.Create(&game)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to create response",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    game,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
