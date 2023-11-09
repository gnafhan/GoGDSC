package credential

import (
	"GoGDSC/data"
	"GoGDSC/middleware"
	models "GoGDSC/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var thisUser models.User
	if err := c.BindJSON(&thisUser); err != nil {
		return
	}
	for i := 1; i < len(data.Users)+1; i++ {
		if thisUser.Username == data.Users[i-1].Username {
			if thisUser.Password == data.Users[i-1].Password {
				thisUser.UserId = data.Users[i-1].UserId
				// create token
				token, err := middleware.CreateJWT(thisUser.Username)
				if err != nil {
					c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
						Message:    "Bad Request",
						StatusCode: 403,
						Data:       []interface{}{},
					})
					return
				}
				c.IndentedJSON(http.StatusOK, gin.H{"token": token})

				break
			} else {
				fmt.Println("cek masuk")
				c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
					Message:    "Bad Request",
					StatusCode: 403,
					Data:       []interface{}{},
				})
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, models.ErrorBadRequest{
				Message:    "Bad Request",
				StatusCode: 403,
				Data:       []interface{}{},
			})
		}
	}
}
