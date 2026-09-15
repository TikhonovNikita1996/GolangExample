package model

type OrderPaid struct {
	EventUuid       string
	OrderUuid       string
	UserUuid        string
	PaymentMethod   string
	TransactionUuid string
}

type ShipAssembled struct {
	EventUuid    string
	OrderUuid    string
	UserUuid     string
	BuildTimeSec int64
}
