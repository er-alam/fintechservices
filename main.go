package main

import (
	controller "fintechservices/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name     string    `json:"name" binding:"required"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday" binding:"required", time_format:"02-01-2006"`
}

func main() {
	gin.DisableConsoleColor()

	var file_name string = "gin.log"
	logfile, err := os.Create(file_name)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	if err != nil {
		panic(err)
	}

	gin.DefaultWriter = io.MultiWriter(logfile)

	route := gin.Default()
	route.GET("/index", controller.Index)
	route.POST("/test", func(c *gin.Context) {
		var body Person

		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		poit := &body
		logfile.WriteString(fmt.Sprintf("%+v\n", *poit))
		c.JSON(http.StatusAccepted, &body)
	})
	route.Run()
}
