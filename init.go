package main

import (
	"GoGDSC/api/todo"
	"GoGDSC/credential"
	"GoGDSC/middleware"
	"context"

	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func postFirebase(c *gin.Context) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./secret.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}
	defer client.Close()

	_, _, err = client.Collection("todos").Add(ctx, map[string]interface{}{
		"name":  "John Doe",
		"email": "johndoe@example.com",
	})
	if err != nil {
		log.Fatalf("Error adding document: %v", err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/", middleware.Validate, todo.GetTodos)
	router.GET("/:id", middleware.Validate, todo.GetTodoById)
	router.POST("/firebase", postFirebase)
	router.POST("/", middleware.Validate, todo.PostTodos)
	router.POST("/login", credential.Login)
	router.PUT("/:id", middleware.Validate, todo.EditTodo)
	router.DELETE("/:id", middleware.Validate, todo.DeleteById)
	err := router.Run("localhost:1234")
	if err != nil {
		return
	}
}
