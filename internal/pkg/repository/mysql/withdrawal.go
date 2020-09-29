package mysql

import (
	"context"

	"gitlab.silkrode.com.tw/golang/kbc_pkg/logger"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/utils/errors"
	"gorm.io/gorm"
)

func NewWithdrawalRepository(db *gorm.DB, log *logger.Logger) domain.WithdrawalRepository {
	return &withdrawalRepository{
		db:  db,
		log: log,
	}
}

type withdrawalRepository struct {
	db  *gorm.DB
	log *logger.Logger
}

func (w *withdrawalRepository) CreateMainOrder(ctx context.Context, mainOrder *domain.MainOrder) (*domain.MainOrder, error) {
	if err := w.db.Table("orders").Create(&mainOrder).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Create main order Err %s", err.Error())
	}
	return mainOrder, nil
}

func (w *withdrawalRepository) CreateSubOrder(ctx context.Context, subOrder *domain.SubOrder) (*domain.SubOrder, error) {
	if err := w.db.Table("sub_order_histories").Create(&subOrder).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Create sub order Err %s", err.Error())
	}
	return subOrder, nil
}

func (w *withdrawalRepository) CreateOrderHistory(ctx context.Context, orderHistory *domain.OrderHistory) (*domain.OrderHistory, error) {
	if err := w.db.Table("order_histories").Create(&orderHistory).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Create Order History Err %s", err.Error())
	}
	return orderHistory, nil
}

func (w *withdrawalRepository) UpdateMainOrder(ctx context.Context, mainOrder *domain.MainOrder) (*domain.MainOrder, error) {
	if err := w.db.Table("orders").Where("tracking_number = ?", mainOrder.TrackingNumber).Updates(&mainOrder).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Update Main Order Err %s", err.Error())
	}
	return mainOrder, nil
}

func (w *withdrawalRepository) UpdateSubOrder(ctx context.Context, subOrder *domain.SubOrder) (*domain.SubOrder, error) {
	if err := w.db.Table("sub_order_histories").Where("sub_tracking_number = ?", subOrder.SubTrackingNumber).Updates(&subOrder).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Update Main Order Err %s", err.Error())
	}
	return subOrder, nil
}

func (w *withdrawalRepository) QueryMainOrder(ctx context.Context, trackingNumber string) (*domain.MainOrder, error) {
	mainOrder := domain.MainOrder{}
	if err := w.db.Table("orders").Find(&mainOrder).Where("tracking_number = ?", trackingNumber).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Query Main Order Err %s", err.Error())
	}
	return &mainOrder, nil
}

func (w *withdrawalRepository) QuerySubOrder(ctx context.Context, subTrackingNumber string) (*domain.SubOrder, error) {
	subOrder := domain.SubOrder{}
	if err := w.db.Table("sub_order_histories").Find(&subOrder).Where("sub_tracking_number = ?", subTrackingNumber).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Query Sub Order Err %s", err.Error())
	}
	return &subOrder, nil
}

func (w *withdrawalRepository) QueryOrderHistory(ctx context.Context, trackingNumber string) (*domain.OrderHistory, error) {
	orderHistory := domain.OrderHistory{}
	if err := w.db.Table("order_histories").Find(&orderHistory).Where("tracking_number = ?", trackingNumber).Error; err != nil {
		return nil, errors.Wrapf(domain.ErrDbOperation, "Query Order History Err %s", err.Error())
	}
	return &orderHistory, nil
}
