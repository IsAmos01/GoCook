package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/upload", Upload)
	
	_ = r.Run(":8080")
}
