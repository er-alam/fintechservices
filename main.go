package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
	Id   int    `json:"lead_id" binding:"required"`
}

func main() {

	file_name := time.Now().Format("2006-01-02 15:04:05") + "_gin.log"
	f, err := os.Create("logs/" + file_name)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		var person Person
		err := c.BindJSON(&person)
		// req := fmt.Sprintf("Received request: %s\n", )
		f.WriteString(fmt.Sprintf("Request : %+v\n", person))

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, person)

	})
	r.Run(":8080")
}
