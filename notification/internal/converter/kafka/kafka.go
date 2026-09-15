package kafka

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/model"
)

type OrderPaidDecoder interface {
	Decode(data []byte) (model.OrderPaid, error)
}

type ShipAssembledDecoder interface {
	Decode(data []byte) (model.ShipAssembled, error)
}
