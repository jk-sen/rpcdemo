package services

import (
    "context"
)

type OrderService struct {
}

func (this *OrderService) CreateOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
    return &OrderResponse{Status: 1, Message: "ok"}, nil
}
