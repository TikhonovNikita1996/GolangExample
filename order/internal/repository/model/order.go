package model

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartsUUID       []string
	TotalPrice      float32
	TransactionUUID *string
	PaymentMethod   *string
	Status          Status
}

type CreateOrderRequest struct {
	OrderUUID  string
	UserUUID   string
	PartUuids  []string
	TotalPrice float32
}

type CreateOrderResponse struct {
	OrderUUID  string
	TotalPrice float32
}

type Status string

const (
	Status_Paid              Status = "Paid"
	Status_Canceled          Status = "Canceled"
	Status_WaitingForPayment Status = "WaitingForPayment"
	Status_Assemdled         Status = "ASSEMBLED"
)
