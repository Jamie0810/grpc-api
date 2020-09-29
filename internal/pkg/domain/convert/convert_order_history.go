package convert

import (
	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain"
)

// OrderHistory converts structs of order history between Prototype and DB model
func OrderHistory(isFromProto bool, pbOrderHistory *pb.OrderHistory, dbOrderHistory *domain.OrderHistory) (*pb.OrderHistory, *domain.OrderHistory) {
	if isFromProto {
		dbOrderHistory = &domain.OrderHistory{
			TrackingNumber:  pbOrderHistory.TrackingNumber,
			ChannelName:     pbOrderHistory.ChannelName,
			IsSuccess:       pbOrderHistory.IsSuccess,
			ResponseMessage: pbOrderHistory.ResponseMessage,
		}
		return nil, dbOrderHistory
	} else {
		pbOrderHistory := &pb.OrderHistory{
			TrackingNumber:  dbOrderHistory.TrackingNumber,
			ChannelName:     dbOrderHistory.ChannelName,
			IsSuccess:       dbOrderHistory.IsSuccess,
			ResponseMessage: dbOrderHistory.ResponseMessage,
		}
		return pbOrderHistory, nil
	}
}
