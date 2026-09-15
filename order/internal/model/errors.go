package model

import "errors"

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderCanNotBePayed = errors.New("order can not be payed")
	ErrPayment            = errors.New("payment was not successful")
	ErrPartsNotAvailable  = errors.New("list of parts is empty")
	ErrPartNotFound       = errors.New("part not found")
	ErrExternal           = errors.New("external error")
)
