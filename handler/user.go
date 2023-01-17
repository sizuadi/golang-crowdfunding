package handler

import (
	"golang-crowdfunding/helper"
	"golang-crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessages := gin.H{"errors": errors}
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "error", errorMessages)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentoken")
	response := helper.ApiResponse("Your account has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
