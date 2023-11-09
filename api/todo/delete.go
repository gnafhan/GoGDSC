package todo

import (
	"GoGDSC/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	index := -1
	x, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	for i, obj := range data.Todos {
		if obj.Id == x {
			index = i
			break
		}
	}
	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
		return
	}

	data.Todos = append(data.Todos[:index], data.Todos[index+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "object deleted"})
}
