package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employee Employee

	if error := c.BindJSON(&employee); error != nil {
		fmt.Printf("failed to bind employee %s\n", error.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: error.Error(),
		})

		return
	}

	h.storage.Insert(&employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	var employee Employee

	if error != nil {
		fmt.Printf("failed to convert param to int %s\n", error.Error())

		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: error.Error(),
		})

		return
	}

	if error := c.BindJSON(&employee); error != nil {
		fmt.Printf("failed to bind employee %s\n", error.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: error.Error(),
		})
	}

	h.storage.Update(id, employee)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employee.ID,
	})
}

func (h *Handler) getEmployee(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		fmt.Printf("failed to convert param to int %s\n", error.Error())

		c.JSON(http.StatusNotFound, ErrorResponse{
			Message: error.Error(),
		})
	}

	empoyee, error := h.storage.Get(id)

	if error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Message: error.Error(),
		})
	}

	c.JSON(http.StatusOK, empoyee)
}

func (h *Handler) RemoveEmployee(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))

	if error != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Message: error.Error(),
		})
	}

	h.storage.Delete(id)

	c.JSON(http.StatusOK, "employee removed")
}
