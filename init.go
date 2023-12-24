package main

import (
	"GoGDSC/api/folder_firebase"
	"GoGDSC/api/todo"
	todofirebase "GoGDSC/api/todo_firebase"
	"GoGDSC/credential"
	"GoGDSC/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	router.Use(cors.New(config))
	router.Use(middleware.CorsMiddleware())
	router.GET("/", todo.GetTodos)
	router.GET("/:id", middleware.Validate, todo.GetTodoById)
	router.GET("/firebase/:folderId", middleware.Validate, todofirebase.GetFirebase)
	router.GET("/firebase/:folderId/:id", middleware.Validate, todofirebase.GetFirebaseById)
	router.POST("/firebase", middleware.Validate, todofirebase.PostFirebase)
	router.POST("/", middleware.Validate, todo.PostTodos)
	router.POST("/login", credential.LoginFirebase)
	router.POST("/register", credential.RegisterFirebase)
	router.PUT("/:id", middleware.Validate, todo.EditTodo)
	router.PUT("/firebase/:folderId/:id", middleware.Validate, todofirebase.EditFireBaseById)
	router.DELETE("/:id", middleware.Validate, todo.DeleteById)
	router.GET("/firebase/folder/", middleware.Validate, folder_firebase.GetFolderFirebase)
	router.DELETE("/firebase/:folderId/:id", middleware.Validate, todofirebase.DeleteFirebaseById)
	router.POST("/firebase/folder/", middleware.Validate, folder_firebase.PostFolderFirebase)
	router.GET("/firebase/folder/:id", middleware.Validate, folder_firebase.GetFolderFirebaseById)
	router.DELETE("/firebase/folder/:id", middleware.Validate, folder_firebase.DeleteFolderFirebaseById)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Route not founds"})
	})
	err := router.Run(":1234")
	if err != nil {
		return
	}
}
