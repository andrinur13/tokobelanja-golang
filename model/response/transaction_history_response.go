package response

type NewTransactionResponse struct {
	Message         string      `json:"message"`
	TransactionBill interface{} `json:"transaction_bill"`
}

type NewTransactionBillResponse struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}
