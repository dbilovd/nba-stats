package main

import (
	"takehome/database"
	"takehome/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	flag.Parse()

	if err := database.Open(); err != nil {
		panic("Failed to connect to Database")
	}

	defer database.Close()

	database.RunMigrations()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
		Views:   html.New("./views", ".html"),
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	api := app.Group("/api")
	api.Get("/player-stats", handlers.GetPlayerStats)
	api.Get("/team-stats", handlers.GetTeamStats)
	api.Post("/player-stats", handlers.PostPlayerStats)
	api.Post("/teams", handlers.PostTeam)
	api.Post("/players", handlers.PostPlayer)
	api.Post("/players/:playerId/teams", handlers.PostPlayerTeamHistory)
	api.Post("/seasons", handlers.PostPlayer)
	api.Post("/games", handlers.PostGames)

	app.Get("/", handlers.GetReports)
	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(*port))
}
