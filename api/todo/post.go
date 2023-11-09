package todo

import (
	"GoGDSC/data"
	models "GoGDSC/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTodos(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	idMax := 0

	for _, a := range data.Todos {
		if a.Id > idMax {
			idMax = a.Id
		}
	}
	newTodo.Id = idMax + 1
	data.Todos = append(data.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
