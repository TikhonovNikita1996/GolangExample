package kafka

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/assembly/internal/model"
)

type OrderPaidDecoder interface {
	Decode(data []byte) (model.OrderPaid, error)
}
