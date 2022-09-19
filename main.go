package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	storage := NewMemoryStorage()
	handler := NewHandler(storage)

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.getEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.RemoveEmployee)

	router.Run()
}
