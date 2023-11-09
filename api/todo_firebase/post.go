package todofirebase

import (
	connect_firebase "GoGDSC/firebase"
	"GoGDSC/middleware"
	models "GoGDSC/model"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func PostFirebase(c *gin.Context) {
	var newTodo models.TodoFirebase
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	client := connect_firebase.Connection()
	defer client.Close()

	payload := middleware.DecodeToken(c)
	if payload == nil {
		fmt.Println("Payload is nil")
		c.IndentedJSON(400, gin.H{"message": "Payload is nil"})
		return
	}

	_, _, err := client.Collection("todos").Add(context.Background(), map[string]interface{}{
		"userId":      payload["userId"],
		"title":       newTodo.Title,
		"description": newTodo.Description,
		"startDate":   newTodo.StartDate,
		"status":      newTodo.Status,
	})
	if err != nil {
		log.Fatalf("Error adding document: %v", err)
	}

	c.IndentedJSON(200, gin.H{"message": "success"})
	return
}
