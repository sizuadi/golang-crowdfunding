package handler

import (
	"golang-crowdfunding/helper"
	"golang-crowdfunding/transaction"
	"golang-crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransationsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		errors := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign's Transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionByUserID(userID)
	if err != nil {
		errors := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to get user's tansactions", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("User's Transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ApiResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		errors := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Create Transaction", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}
