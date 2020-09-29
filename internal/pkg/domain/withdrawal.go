package domain

import (
	"context"

	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
)

type WithdrawalUsecase interface {
	pb.OrderServer
}

type WithdrawalRepository interface {
	CreateMainOrder(ctx context.Context, mainOrder *MainOrder) (*MainOrder, error)
	CreateSubOrder(ctx context.Context, subOrder *SubOrder) (*SubOrder, error)
	CreateOrderHistory(ctx context.Context, orderHistory *OrderHistory) (*OrderHistory, error)
	UpdateMainOrder(ctx context.Context, mainOrder *MainOrder) (*MainOrder, error)
	UpdateSubOrder(ctx context.Context, subOrder *SubOrder) (*SubOrder, error)
	QueryMainOrder(ctx context.Context, trackingNumber string) (*MainOrder, error)
	QuerySubOrder(ctx context.Context, subTrackingNumber string) (*SubOrder, error)
	QueryOrderHistory(ctx context.Context, trackingNumber string) (*OrderHistory, error)
}
