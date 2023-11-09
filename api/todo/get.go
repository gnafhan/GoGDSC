package todo

import (
	"GoGDSC/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	for _, a := range data.Todos {
		if i == a.Id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
