package controller

import (
	"net/http"
	"tokobelanja-golang/helper"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/model/response"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
)

type transactionHistoryController struct {
	transactionHistoryService service.TransactionHistoryService
	userService               service.UserService
}

func NewTransactionHistoryController(transactionHistoryService service.TransactionHistoryService, userService service.UserService) *transactionHistoryController {
	return &transactionHistoryController{transactionHistoryService, userService}
}

func (h *transactionHistoryController) NewTransaction(c *gin.Context) {
	var input input.InputTransaction

	err := c.ShouldBindJSON(&input)

	if err != nil {
		resp := helper.APIResponse("error", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newTransaction, err := h.transactionHistoryService.CreateTransaction(input)

	if err != nil {
		resp := helper.APIResponse("error", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	billResponse := response.NewTransactionBillResponse{
		TotalPrice:   newTransaction.TotalPrice,
		Quantity:     input.Quantity,
		ProductTitle: "nunggu product domain",
	}

	newTransactionResponse := response.NewTransactionResponse{
		Message:         "You have successfully purchased the product",
		TransactionBill: billResponse,
	}

	resp := helper.APIResponse("success", newTransactionResponse)
	c.JSON(http.StatusCreated, resp)
}

func (h *transactionHistoryController) GetMyTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBind(&currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactions, err := h.transactionHistoryService.GetMyTransaction(currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("success", transactions)
	c.JSON(http.StatusOK, response)
	return
}

func (h *transactionHistoryController) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBind(&currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userdata, err := h.userService.CheckUserAdmin(currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if userdata == false {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
	}

	h.transactionHistoryService.GetMyTransaction(2)

	response := helper.APIResponse("success", "success get user transactions")
	c.JSON(http.StatusUnauthorized, response)

}
