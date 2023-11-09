package todofirebase

import (
	connect_firebase "GoGDSC/firebase"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetFirebase(c *gin.Context) {
	client := connect_firebase.Connection()
	defer client.Close()

	datas := []map[string]interface{}{}

	iter := client.Collection("todos").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		data := doc.Data()
		data["id"] = doc.Ref.ID
		datas = append(datas, data)
	}

	payload := gin.H{
		"message": "success",
		"data":    datas,
	}

	c.IndentedJSON(200, payload)
	return

}
