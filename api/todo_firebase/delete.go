package todofirebase

import (
	connect_firebase "GoGDSC/firebase"
	"GoGDSC/middleware"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func DeleteFirebaseById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	payload := middleware.DecodeToken(c)
	if payload == nil {
		fmt.Println("Payload is nil")
		c.IndentedJSON(400, gin.H{"message": "Payload is nil"})
		return
	}

	docs, err := client.Collection("todos").Where("userId", "==", payload["userId"]).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("Failed to get documents: %v", err)
	}

	for _, doc := range docs {
		if id == doc.Ref.ID {
			_, err := client.Collection("todos").Doc(id).Delete(context.Background())
			if err != nil {
				log.Fatalf("Failed adding document: %v", err)
			}
			c.IndentedJSON(200, gin.H{"message": "success"})
			return
		} else {
			c.IndentedJSON(400, gin.H{"message": "id not found"})
			return
		}
	}
	return
}
