package index

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `json:"name" binding:"required"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday" binding:"required", time_format:"02-01-2006"`
}

func Index(c *gin.Context) {

	var person Person
	err := c.BindJSON(&person)
	if err == nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday,
		"message":  "success",
	})
}
