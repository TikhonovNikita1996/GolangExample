package v1

import (
	paymentV1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

type paymentServiceClient struct {
	generatedClient paymentV1.PaymentServiceClient
}

func NewClient(generatedClient paymentV1.PaymentServiceClient) *paymentServiceClient {
	return &paymentServiceClient{generatedClient}
}
