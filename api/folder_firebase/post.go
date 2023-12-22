package folder_firebase

import (
	connect_firebase "GoGDSC/firebase"
	"GoGDSC/middleware"
	models "GoGDSC/model"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func PostFolderFirebase(c *gin.Context) {
	var newFolder models.Folder
	if err := c.BindJSON(&newFolder); err != nil {
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

	_, _, err := client.Collection("folders").Add(context.Background(), map[string]interface{}{
		"userId": payload["userId"],
		"name":   newFolder.Name,
	})
	if err != nil {
		log.Fatalf("Error adding document: %v", err)
	}

	c.IndentedJSON(200, gin.H{"message": "success"})
	return
}
