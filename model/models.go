package models

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"startDate"`
}
type ErrorBadRequest struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type User struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}
