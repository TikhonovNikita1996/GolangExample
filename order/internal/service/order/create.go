package order

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *OrderService) CreateOrder(ctx context.Context, req *model.CreateOrderRequest) (model.CreateOrderResponse, error) {
	filter := model.PartsFilter{
		Uuids: req.PartUuids,
	}
	parts, err := s.inventoryClient.GetParts(ctx, filter)
	if err != nil {
		return model.CreateOrderResponse{}, fmt.Errorf("inventoryClient.GetParts: %w", err)
	}
	if len(parts) == 0 {
		return model.CreateOrderResponse{}, fmt.Errorf("inventoryClient.GetParts: %w", err)
	}

	notAvailableParts, err := checkIfAllPartsAreExist(parts, req.PartUuids)
	if len(notAvailableParts) == 0 {
		totalPrice := 0.0
		for _, part := range parts {
			totalPrice += part.Price
		}
		req.TotalPrice = float32(totalPrice)

		req.OrderUUID = uuid.NewString()
		orderCreationResponse, err := s.orderRepository.CreateOrder(ctx, req)
		if err != nil {
			return model.CreateOrderResponse{}, fmt.Errorf("orderRepository.CreateOrder: %w", err)
		}

		return model.CreateOrderResponse{
			OrderUUID:  orderCreationResponse,
			TotalPrice: float32(totalPrice),
		}, nil
	}
	return model.CreateOrderResponse{}, fmt.Errorf("parts not available: %w", model.ErrPartsNotAvailable)
}

func checkIfAllPartsAreExist(parts []*model.Part, partsUuids []string) ([]string, error) {
	var missing []string
	if len(parts) == len(parts) {
		for _, s2 := range partsUuids {
			found := false
			for _, s1 := range parts {
				if s2 == s1.UUID {
					found = true
					break
				}
			}
			if !found {
				missing = append(missing, s2)
			}
		}
		return missing, nil
	}
	return nil, model.ErrPartsNotAvailable
}
