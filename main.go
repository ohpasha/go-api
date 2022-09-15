package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	storage := NewMemoryStorage()
	handler := NewHandler(storage)

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id")
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id")

	router.Run()
}
