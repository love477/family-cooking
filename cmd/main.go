package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
}

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Message: "hello",
	})
}

func main() {
	r := gin.Default()

	r.GET("/hello", Hello)

	r.Run(":8080")
}
