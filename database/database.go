package database

import (
	"fmt"
	"takehome/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() error {
	db, err := gorm.Open(sqlite.Open("stats.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	DB = db

	if err != nil {
		// panic("Failed to connect to DB")
		return err
	} else {
		fmt.Println("Connected with Database")
	}

	return nil
}

func Close() {
	// return DB.Close()
}

func RunMigrations() {
	DB.AutoMigrate(
		&models.Player{},
		&models.Team{},
		&models.PlayerTeamHistory{},
		&models.Season{},
		&models.Game{},
		&models.PlayerGameStat{},
	)
}
