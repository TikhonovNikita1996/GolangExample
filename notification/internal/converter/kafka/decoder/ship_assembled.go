package decoder

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/model"
	eventsv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/events/v1"
)

type shipAssembledDecoder struct{}

func NewShipAssembledDecoder() *shipAssembledDecoder {
	return &shipAssembledDecoder{}
}

func (d *shipAssembledDecoder) Decode(data []byte) (model.ShipAssembled, error) {
	var pb eventsv1.ShipAssembled
	if err := proto.Unmarshal(data, &pb); err != nil {
		return model.ShipAssembled{}, fmt.Errorf("failed to unmarshal protobuf: %w", err)
	}

	return model.ShipAssembled{
		EventUuid:    pb.EventUuid,
		OrderUuid:    pb.OrderUuid,
		UserUuid:     pb.UserUuid,
		BuildTimeSec: pb.BuildTimeSec,
	}, nil
}
