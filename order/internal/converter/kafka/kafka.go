package kafka

import "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"

type ShipAssembledDecoder interface {
	Decode(data []byte) (model.ShipAssembled, error)
}
