package usecase

import (
	"context"

	"github.com/rs/xid"
	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
	"gitlab.silkrode.com.tw/golang/kbc_pkg/logger"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain/convert"
)

type withdrawalUsecase struct {
	withdrawalRepo domain.WithdrawalRepository
	log            *logger.Logger
}

func NewWithdrawalUsecase(w domain.WithdrawalRepository, log *logger.Logger) domain.WithdrawalUsecase {
	return &withdrawalUsecase{
		log:            log,
		withdrawalRepo: w,
	}
}

const (
	errorKeyPrefix = "grpc_create_order_"
)

func (w *withdrawalUsecase) CreateMainOrder(ctx context.Context, req *pb.MainOrder) (resp *pb.MainOrder, err error) {
	trackingNumber := xid.New().String()

	_, dbMainOrder := convert.MainOrder(true, trackingNumber, req, nil)
	dbMainOrder, err = w.withdrawalRepo.CreateMainOrder(ctx, dbMainOrder)
	if err != nil {
		errorKey := errorKeyPrefix + dbMainOrder.TrackingNumber
		w.log.Error().Str("CreateMainOrder Error Key", errorKey).Str("Create main order Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.MainOrder(false, trackingNumber, nil, dbMainOrder)
	return resp, nil
}

func (w *withdrawalUsecase) CreateSubOrder(ctx context.Context, req *pb.SubOrder) (resp *pb.SubOrder, err error) {
	_, dbSubOrder := convert.SubOrder(true, req, nil)
	dbSubOrder, err = w.withdrawalRepo.CreateSubOrder(ctx, dbSubOrder)
	if err != nil {
		errorKey := errorKeyPrefix + dbSubOrder.TrackingNumber
		w.log.Error().Str("CreateSubOrder Error Key", errorKey).Str("Create Sub order Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.SubOrder(false, nil, dbSubOrder)
	return resp, nil
}

func (w *withdrawalUsecase) CreateOrderHistory(ctx context.Context, req *pb.OrderHistory) (resp *pb.OrderHistory, err error) {
	_, dbOrderHistory := convert.OrderHistory(true, req, nil)
	dbOrderHistory, err = w.withdrawalRepo.CreateOrderHistory(ctx, dbOrderHistory)
	if err != nil {
		errorKey := errorKeyPrefix + dbOrderHistory.TrackingNumber
		w.log.Error().Str("CreateOrderHistory Error Key", errorKey).Str("Create Order History Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.OrderHistory(false, nil, dbOrderHistory)
	return resp, nil
}

func (w *withdrawalUsecase) UpdateMainOrder(ctx context.Context, req *pb.MainOrder) (resp *pb.MainOrder, err error) {
	_, dbMainOrder := convert.MainOrder(true, req.TrackingNumber, req, nil)
	dbMainOrder, err = w.withdrawalRepo.UpdateMainOrder(ctx, dbMainOrder)
	if err != nil {
		errorKey := errorKeyPrefix + dbMainOrder.TrackingNumber
		w.log.Error().Str("UpdateMainOrder Error Key", errorKey).Str("Update main order Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.MainOrder(false, req.TrackingNumber, nil, dbMainOrder)
	return resp, nil
}

func (w *withdrawalUsecase) UpdateSubOrder(ctx context.Context, req *pb.SubOrder) (resp *pb.SubOrder, err error) {
	_, dbSubOrder := convert.SubOrder(true, req, nil)
	dbSubOrder, err = w.withdrawalRepo.UpdateSubOrder(ctx, dbSubOrder)
	if err != nil {
		errorKey := errorKeyPrefix + dbSubOrder.SubTrackingNumber
		w.log.Error().Str("UpdateMainOrder Error Key", errorKey).Str("Update sub order Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.SubOrder(false, nil, dbSubOrder)
	return resp, nil
}

func (w *withdrawalUsecase) QueryMainOrder(ctx context.Context, trackingNumber *pb.TrackingNumber) (resp *pb.MainOrder, err error) {
	mainOrder, err := w.withdrawalRepo.QueryMainOrder(ctx, trackingNumber.TrackingNumber)
	if err != nil {
		errorKey := errorKeyPrefix + trackingNumber.TrackingNumber
		w.log.Error().Str("QueryMainOrder Error Key", errorKey).Str("Query main order Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.MainOrder(false, trackingNumber.TrackingNumber, nil, mainOrder)
	return resp, nil
}

func (w *withdrawalUsecase) QuerySubOrder(ctx context.Context, subTrackingNumber *pb.SubTrackingNumber) (resp *pb.SubOrder, err error) {
	subOrder, err := w.withdrawalRepo.QuerySubOrder(ctx, subTrackingNumber.SubTrackingNumber)
	if err != nil {
		errorKey := errorKeyPrefix + subTrackingNumber.SubTrackingNumber
		w.log.Error().Str("QuerySubOrder Error Key", errorKey).Str("Query sub order Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.SubOrder(false, nil, subOrder)
	return resp, nil
}

func (w *withdrawalUsecase) QueryOrderHistory(ctx context.Context, trackingNumber *pb.TrackingNumber) (resp *pb.OrderHistory, err error) {
	orderHistory, err := w.withdrawalRepo.QueryOrderHistory(ctx, trackingNumber.TrackingNumber)
	if err != nil {
		errorKey := errorKeyPrefix + trackingNumber.TrackingNumber
		w.log.Error().Str("QueryOrderHistory Error Key", errorKey).Str("Query order history Error", err.Error()).Send()
		return nil, err
	}

	resp, _ = convert.OrderHistory(false, nil, orderHistory)
	return resp, nil
}
