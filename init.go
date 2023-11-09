package main

import (
	"GoGDSC/api/todo"
	todofirebase "GoGDSC/api/todo_firebase"
	"GoGDSC/credential"
	"GoGDSC/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", middleware.Validate, todo.GetTodos)
	router.GET("/:id", middleware.Validate, todo.GetTodoById)
	router.GET("/firebase", middleware.Validate, todofirebase.GetFirebase)
	router.GET("/firebase/:id", middleware.Validate, todofirebase.GetFirebaseById)
	router.POST("/firebase", middleware.Validate, todofirebase.PostFirebase)
	router.POST("/", middleware.Validate, todo.PostTodos)
	router.POST("/login", credential.LoginFirebase)
	router.POST("/register", credential.RegisterFirebase)
	router.PUT("/:id", middleware.Validate, todo.EditTodo)
	router.PUT("/firebase/:id", middleware.Validate, todofirebase.EditFireBaseById)
	router.DELETE("/:id", middleware.Validate, todo.DeleteById)
	router.DELETE("/firebase/:id", middleware.Validate, todofirebase.DeleteFirebaseById)
	err := router.Run("localhost:1234")
	if err != nil {
		return
	}
}
