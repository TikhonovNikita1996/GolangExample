package v1

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	paymentV1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

func (api *Api) PayOrder(ctx context.Context, req *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	result, err := api.paymentService.PayOrder(ctx, req.Info.GetOrderUuid(), req.Info.GetUserUuid())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}
	return &paymentV1.PayOrderResponse{
		TransactionUuid: result,
	}, nil
}
