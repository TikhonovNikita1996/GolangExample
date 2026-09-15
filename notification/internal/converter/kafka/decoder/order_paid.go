package decoder

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/model"
	eventsv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/events/v1"
)

type orderPaidDecoder struct{}

func NewOrderPaidDecoder() *orderPaidDecoder {
	return &orderPaidDecoder{}
}

func (d *orderPaidDecoder) Decode(data []byte) (model.OrderPaid, error) {
	var pb eventsv1.OrderPayed
	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.OrderPaid{}, fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return model.OrderPaid{
		EventUuid:       pb.EventUuid,
		OrderUuid:       pb.OrderUuid,
		UserUuid:        pb.UserUuid,
		PaymentMethod:   pb.PaymentMethod,
		TransactionUuid: pb.TransactionUuid,
	}, nil
}
