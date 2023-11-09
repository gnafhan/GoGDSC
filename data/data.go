package data

import (
	models "GoGDSC/model"
	"time"
)

var Todos = []models.Todo{
	{
		Id:          1,
		Title:       "Mencuci pakaian",
		Description: "mencuci pakaian ",
		Status:      "on going",
		StartDate:   time.Date(2023, time.November, 15, 23, 0, 0, 0, time.FixedZone("WIB", 7*60*60)),
	},
	{
		Id:          2,
		Title:       "Belajar",
		Description: "Belajar",
		Status:      "done",
		StartDate:   time.Date(2023, time.November, 3, 23, 0, 0, 0, time.FixedZone("WIB", 7*60*60)),
	},
	{
		Id:          3,
		Title:       "Mobile legend",
		Description: "main cut",
		Status:      "soon",
		StartDate:   time.Date(2023, time.December, 15, 23, 0, 0, 0, time.FixedZone("WIB", 7*60*60)),
	},
}

var Users = []models.User{
	{
		UserId:   1,
		Username: "nafhan",
		Password: "12345",
	},
	{
		UserId:   2,
		Username: "joko",
		Password: "tidur",
	},
	{
		UserId:   3,
		Username: "Asep",
		Password: "bangun",
	},
}
