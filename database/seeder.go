package database

import (
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	// // Teams
	// teams := []models.Team{
	// 	{Name: "City Fires"},
	// 	{Name: "Town Waters"},
	// }

	// team_1_players = []models.Player{
	// 	{Name: "First City"},
	// 	{Name: "Second City"},
	// 	{Name: "Third City"},
	// 	{Name: "Fourt City"},
	// }

	// team_2_players = []models.Player{
	// 	{Name: "First Towns"},
	// 	{Name: "Second Towns"},
	// 	{Name: "Third Towns"},
	// 	{Name: "Fourt Towns"},
	// }

	// date_layout := "2000-01-01"
	// start_date, _ := time.Parse(date_layout, "2023-10-01")
	// end_date, _ := time.Parse(date_layout, "2024-06-30")
	// seasons = []models.Season{
	// 	{
	// 		Name:      "2023/24",
	// 		StartedAt: datatypes.Date(start_date),
	// 		EndedAt:   datatypes.Date(end_date),
	// 	},
	// }

	// // Create records in database
	// for _, team := range teams {
	// 	err := db.Save(&team).Error
	// 	if err != nil {

	// 	}

	// }
}
