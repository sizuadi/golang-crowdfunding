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