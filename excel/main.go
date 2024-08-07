package main

import "github.com/gin-gonic/gin"

func main() {
	c := gin.Default()
	c.POST("/upload", Upload)
	
	_ = c.Run(":8080")
	
}
