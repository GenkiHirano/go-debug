package main

import (
	"github.com/GenkiHirano/go-debug/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/product", handler.CreateProduct())

	r.Run()
}
