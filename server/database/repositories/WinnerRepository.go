package repositories

import "github.com/piru72/winners-crud/server/database/entities"

var winners = []entities.Winner{
	{ID: "1", Season: "2024", Game: "Chess", Position: "Champion", TeamMember1: "Alice", TeamMember2: "Bob"},
	{ID: "2", Season: "2024", Game: "UNO", Position: "1st Runners Up", TeamMember1: "Bob", TeamMember2: "Charlie"},
	{ID: "3", Season: "2025", Game: "FoosBall", Position: "2nd Runners Up", TeamMember1: "Charlie", TeamMember2: "Dave"},
}

func GetWinners() []entities.Winner {
	return winners
}
