package models

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"startDate"`
}

type Folder struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type TodoFirebase struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	FolderId    string    `json:"folderId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Note        string    `json:"note"`
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

type UserFirebase struct {
	UserId    string `json:"userId"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
