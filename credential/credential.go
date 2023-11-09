package credential

import (
	connect_firebase "GoGDSC/firebase"
	"GoGDSC/middleware"
	models "GoGDSC/model"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginFirebase(c *gin.Context) {
	var loginUser models.UserLogin
	if err := c.BindJSON(&loginUser); err != nil {
		return
	}
	client := connect_firebase.Connection()
	defer client.Close()

	if loginUser.Username == "" || loginUser.Password == "" {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Bad Request",
			StatusCode: 400,
			Data:       []interface{}{},
		})
		return
	}

	docs, err := client.Collection("users").Where("username", "==", loginUser.Username).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("Failed to get documents: %v", err)
	}

	if len(docs) == 0 {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Username not found",
			StatusCode: 400,
			Data:       []interface{}{},
		})
		return
	}

	var user models.UserFirebase
	for _, doc := range docs {
		doc.DataTo(&user)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Wrong password",
			StatusCode: 400,
			Data:       []interface{}{},
		})
		return
	}

	token, err := middleware.CreateJWT(user.Username, docs[0].Ref.ID, user.Email, user.FirstName, user.LastName)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Bad Request",
			StatusCode: 403,
			Data:       []interface{}{},
		})
		return
	}
	user.Password = ""

	c.IndentedJSON(http.StatusOK, gin.H{"token": token, "message": "success", "data": user})
	return

}

func RegisterFirebase(c *gin.Context) {
	var newUser models.UserFirebase
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	client := connect_firebase.Connection()
	defer client.Close()

	if newUser.Username == "" || newUser.Password == "" || newUser.Email == "" || newUser.FirstName == "" || newUser.LastName == "" {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Bad Request",
			StatusCode: 400,
			Data:       []interface{}{},
		})
		return
	}

	if len(newUser.Password) < 6 {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Password < 6",
			StatusCode: 400,
			Data:       []interface{}{},
		})
		return
	}

	docs, err := client.Collection("users").Where("username", "==", newUser.Username).Documents(context.Background()).GetAll()
	if err != nil {
		log.Fatalf("Failed to get documents: %v", err)
	}

	if len(docs) > 0 {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Username already exist",
			StatusCode: 400,
			Data:       []interface{}{},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)
	if err != nil {
		// Handle error
		return
	}

	docRef, _, err := client.Collection("users").Add(context.Background(), map[string]interface{}{
		"username":  newUser.Username,
		"password":  newUser.Password,
		"email":     newUser.Email,
		"firstName": newUser.FirstName,
		"lastName":  newUser.LastName,
	})
	if err != nil {
		// Handle error
		return
	}
	token, err := middleware.CreateJWT(newUser.Username, docRef.ID, newUser.Email, newUser.FirstName, newUser.LastName)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
			Message:    "Bad Request",
			StatusCode: 403,
			Data:       []interface{}{},
		})
		return
	}
	newUser.Password = ""

	c.IndentedJSON(http.StatusOK, gin.H{"token": token, "message": "success", "data": newUser})
	return

}
