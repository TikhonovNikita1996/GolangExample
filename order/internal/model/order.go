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

type OrderUpdate struct {
	OrderUUID       string
	UserUUID        *string
	PartsUUID       []*string
	TotalPrice      *float32
	TransactionUUID *string
	PaymentMethod   *string
	Status          *Status
}

type CreateOrderRequest struct {
	OrderUUID  string
	UserUUID   string
	PartUuids  []string
	TotalPrice float32
}

type PaymentRequest struct {
	OrderUuid       string
	PaymentMethod   string
	TransactionUuid string
	PaymentStatus   Status
}

type CreateOrderResponse struct {
	OrderUUID  string
	TotalPrice float32
}

type CancelResponse int32

const (
	Cancel_Response_NoContent CancelResponse = 0
	Cancel_Response_NotFound  CancelResponse = 1
	Cancel_Response_Conflict  CancelResponse = 2
)

type Status string

const (
	Status_Paid              Status = "Paid"
	Status_Canceled          Status = "Canceled"
	Status_WaitingForPayment Status = "WaitingForPayment"
	Status_Assemdled         Status = "ASSEMBLED"
)

type PaymentResponse struct {
	TransactionUuid string
}
