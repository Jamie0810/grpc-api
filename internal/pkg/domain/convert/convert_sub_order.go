package convert

import (
	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain"
)

// SubOrder converts structs of suborder between Prototype and DB model
func SubOrder(isFromProto bool, pbSubOrder *pb.SubOrder, dbSubOrder *domain.SubOrder) (*pb.SubOrder, *domain.SubOrder) {
	if isFromProto {
		dbSubOrder = &domain.SubOrder{
			TrackingNumber:         pbSubOrder.TrackingNumber,
			SubTrackingNumber:      pbSubOrder.SubTrackingNumber,
			ChannelID:              pbSubOrder.ChannelID,
			ChannelName:            pbSubOrder.ChannelName,
			MerchantProductAPIName: pbSubOrder.MerchantProductAPIName,
			IsSuccess:              pbSubOrder.IsSuccess,
			ResponseMessage:        pbSubOrder.ResponseMessage,
			ChannelReq:             pbSubOrder.ChannelReq,
			ChannelResp:            pbSubOrder.ChannelResp,
		}
		return nil, dbSubOrder
	} else {
		pbSubOrder := &pb.SubOrder{
			TrackingNumber:         dbSubOrder.TrackingNumber,
			SubTrackingNumber:      dbSubOrder.SubTrackingNumber,
			ChannelID:              dbSubOrder.ChannelID,
			ChannelName:            dbSubOrder.ChannelName,
			MerchantProductAPIName: dbSubOrder.MerchantProductAPIName,
			IsSuccess:              dbSubOrder.IsSuccess,
			ResponseMessage:        dbSubOrder.ResponseMessage,
			ChannelReq:             dbSubOrder.ChannelReq,
			ChannelResp:            dbSubOrder.ChannelResp,
		}
		return pbSubOrder, nil
	}
}
