package todo

import (
	"GoGDSC/data"
	models "GoGDSC/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func EditTodo(c *gin.Context) {
	var newTodo models.Todo
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	for _, a := range data.Todos {
		if i == a.Id {
			if newTodo.Title != "" {
				a.Title = newTodo.Title
			}
			if newTodo.Description != "" {
				a.Description = newTodo.Description
			}
			if newTodo.StartDate != time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
				a.StartDate = newTodo.StartDate
			}
			if newTodo.Status != "" {
				a.Status = newTodo.Status
			}
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

}
