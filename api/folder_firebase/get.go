package folder_firebase

import (
	connect_firebase "GoGDSC/firebase"
	"GoGDSC/middleware"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func GetFolderFirebase(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	datas := []map[string]interface{}{}

	fmt.Println(c)
	payload := middleware.DecodeToken(c)
	if payload == nil {
		c.IndentedJSON(400, gin.H{"message": "Payload is nil"})
		return
	}

	docs, err := client.Collection("folders").Where("userId", "==", payload["userId"]).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("Failed to get documents: %v", err)
	}

	for _, doc := range docs {
		data := doc.Data()
		data["id"] = doc.Ref.ID
		datas = append(datas, data)
	}

	payloads := gin.H{
		"message": "success",
		"data":    datas,
	}

	c.IndentedJSON(200, payloads)
	return

}

func GetFolderFirebaseById(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	id := c.Param("id")

	payload := middleware.DecodeToken(c)
	if payload == nil {
		c.IndentedJSON(400, gin.H{"message": "Payload is nil"})
		return
	}

	docs, err := client.Collection("folders").Where("userId", "==", payload["userId"]).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("Failed to get documents: %v", err)
	}

	for _, doc := range docs {
		if id == doc.Ref.ID {
			data := doc.Data()
			data["id"] = doc.Ref.ID
			payloads := gin.H{
				"message": "success",
				"data":    data,
			}
			c.IndentedJSON(200, payloads)
			return
		}
	}
	c.IndentedJSON(400, gin.H{"message": "id not found"})
	return
}
