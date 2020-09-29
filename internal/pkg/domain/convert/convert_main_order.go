package convert

import (
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain"
)

// MainOrder converts structs of main order between Prototype and DB model
func MainOrder(isFromProto bool, trackingNumber string, pbMainOrder *pb.MainOrder, dbMainOrder *domain.MainOrder) (*pb.MainOrder, *domain.MainOrder) {
	if isFromProto {
		_, walletStatus := convertWalletStatus(true, pbMainOrder.WalletStatus, "")
		_, paymentType := convertPaymentType(true, pbMainOrder.PaymentType, "")
		_, status := convertStatus(true, pbMainOrder.Status, "")
		_, merchantRateType := convertMerchantRateType(true, pbMainOrder.MerchantRateType, "")
		_, merchantResult := convertMerchantResult(true, pbMainOrder.MerchantResult, "")

		completedAt, err := timeConverter(pbMainOrder.CompletedAt)
		if err != nil {
			log.Logger.Error().Str("timeConverter Err", err.Error())
		}

		deletedAt, err := timeConverter(pbMainOrder.DeletedAt)
		if err != nil {
			log.Logger.Error().Str("timeConverter Err", err.Error())
		}

		expiredAt, err := timeConverter(pbMainOrder.ExpiredAt)
		if err != nil {
			log.Logger.Error().Str("timeConverter Err", err.Error())
		}

		dbMainOrder = &domain.MainOrder{
			TrackingNumber:         trackingNumber,
			WalletStatus:           walletStatus,
			TransactionType:        domain.Withdrawal,
			PaymentNumber:          pbMainOrder.PaymentNumber,
			PaymentType:            paymentType,
			WithdrawAccount:        pbMainOrder.WithdrawAccount,
			Status:                 status,
			PlayerID:               pbMainOrder.PlayerID,
			UserName:               pbMainOrder.UserName,
			DeviceID:               pbMainOrder.DeviceID,
			DeviceIP:               pbMainOrder.DeviceIP,
			DeviceType:             pbMainOrder.DeviceType,
			Telephone:              pbMainOrder.Telephone,
			EstimatedCost:          pbMainOrder.EstimatedCost,
			ActualCost:             pbMainOrder.ActualCost,
			CompletedAt:            *completedAt,
			DeletedAt:              deletedAt,
			ExpiredAt:              expiredAt,
			MerchantID:             pbMainOrder.MerchantID,
			MerchantName:           pbMainOrder.MerchantName,
			MerchantOrderNumber:    pbMainOrder.MerchantOrderNumber,
			MerchantFee:            pbMainOrder.MerchantFee,
			MerchantRateType:       merchantRateType,
			MerchantRate:           pbMainOrder.MerchantRate,
			MerchantRateFixed:      pbMainOrder.MerchantRateFixed,
			MerchantResult:         merchantResult,
			MerchantNotifyURL:      pbMainOrder.MerchantNotifyURL,
			MerchantProjectID:      pbMainOrder.MerchantProjectID,
			MerchantProjectName:    pbMainOrder.MerchantProjectName,
			MerchantProductID:      pbMainOrder.MerchantProductID,
			MerchantProductName:    pbMainOrder.MerchantProductName,
			MerchantProductAPIID:   pbMainOrder.MerchantProductAPIID,
			MerchantProductAPIName: pbMainOrder.MerchantProductAPIName,
			ChannelID:              pbMainOrder.ChannelID,
			ChannelCode:            pbMainOrder.ChannelCode,
			ChannelName:            pbMainOrder.ChannelName,
			ChannelReq:             pbMainOrder.ChannelReq,
			ChannelOrderNumber:     pbMainOrder.ChannelOrderNumber,
			ChannelMerchantID:      pbMainOrder.ChannelMerchantID,
			BBReason:               pbMainOrder.BBReason,
			BBSkip:                 pbMainOrder.BBSkip,
		}
		return nil, dbMainOrder
	} else {
		walletStatus, _ := convertWalletStatus(false, 0, dbMainOrder.WalletStatus)
		paymentType, _ := convertPaymentType(false, 0, dbMainOrder.PaymentType)
		status, _ := convertStatus(false, 0, dbMainOrder.Status)
		merchantRateType, _ := convertMerchantRateType(false, 0, dbMainOrder.MerchantRateType)
		merchantResult, _ := convertMerchantResult(false, 0, dbMainOrder.MerchantResult)

		pbMainOrder := &pb.MainOrder{
			TrackingNumber:  trackingNumber,
			WalletStatus:    walletStatus,
			TransactionType: pb.TransactionType_Withdrawal,
			PaymentNumber:   dbMainOrder.PaymentNumber,
			PaymentType:     paymentType,
			WithdrawAccount: dbMainOrder.WithdrawAccount,
			Status:          status,
			PlayerID:        dbMainOrder.PlayerID,
			UserName:        dbMainOrder.UserName,
			DeviceID:        dbMainOrder.DeviceID,
			DeviceIP:        dbMainOrder.DeviceIP,
			DeviceType:      dbMainOrder.DeviceType,
			Telephone:       dbMainOrder.Telephone,
			EstimatedCost:   dbMainOrder.EstimatedCost,
			ActualCost:      dbMainOrder.ActualCost,
			// CompletedAt
			// DeletedAt
			// ExpiredAt
			MerchantID:             dbMainOrder.MerchantID,
			MerchantName:           dbMainOrder.MerchantName,
			MerchantOrderNumber:    dbMainOrder.MerchantOrderNumber,
			MerchantFee:            dbMainOrder.MerchantFee,
			MerchantRateType:       merchantRateType,
			MerchantRate:           dbMainOrder.MerchantRate,
			MerchantRateFixed:      dbMainOrder.MerchantRateFixed,
			MerchantResult:         merchantResult,
			MerchantNotifyURL:      dbMainOrder.MerchantNotifyURL,
			MerchantProjectID:      dbMainOrder.MerchantProjectID,
			MerchantProjectName:    dbMainOrder.MerchantProjectName,
			MerchantProductID:      dbMainOrder.MerchantProductID,
			MerchantProductName:    dbMainOrder.MerchantProductName,
			MerchantProductAPIID:   dbMainOrder.MerchantProductAPIID,
			MerchantProductAPIName: dbMainOrder.MerchantProductAPIName,
			ChannelID:              dbMainOrder.ChannelID,
			ChannelCode:            dbMainOrder.ChannelCode,
			ChannelName:            dbMainOrder.ChannelName,
			ChannelReq:             dbMainOrder.ChannelReq,
			ChannelOrderNumber:     dbMainOrder.ChannelOrderNumber,
			ChannelMerchantID:      dbMainOrder.ChannelMerchantID,
			BBReason:               dbMainOrder.BBReason,
			BBSkip:                 dbMainOrder.BBSkip,
		}
		return pbMainOrder, nil
	}
}

func convertWalletStatus(isFromProto bool, pbWalletStatus pb.WalletStatus, dbWalletStatus domain.WalletStatus) (pb.WalletStatus, domain.WalletStatus) {
	if isFromProto {
		switch pbWalletStatus {
		case pb.WalletStatus_before:
			return 0, domain.BeforeNotify
		case pb.WalletStatus_confirm:
			return 0, domain.ConfirmNotify
		case pb.WalletStatus_failed:
			return 0, domain.FailedNotify
		}
		return 0, domain.UnknowWalletStatus
	} else {
		switch dbWalletStatus {
		case domain.BeforeNotify:
			return pb.WalletStatus_before, ""
		case domain.ConfirmNotify:
			return pb.WalletStatus_confirm, ""
		case domain.FailedNotify:
			return pb.WalletStatus_failed, ""
		}

		return pb.WalletStatus_unknow_wallet_status, ""
	}
}

func convertPaymentType(isFromProto bool, pbPaymentType pb.PaymentType, dbPaymentType domain.PaymentType) (pb.PaymentType, domain.PaymentType) {
	if isFromProto {
		switch pbPaymentType {
		case pb.PaymentType_ALIPAY:
			return 0, domain.AliPay
		case pb.PaymentType_WECHAT:
			return 0, domain.Wechat
		case pb.PaymentType_BANKCARD:
			return 0, domain.BankCard
		}
		return 0, domain.UnknowPaymentType
	} else {
		switch dbPaymentType {
		case domain.AliPay:
			return pb.PaymentType_ALIPAY, ""
		case domain.Wechat:
			return pb.PaymentType_WECHAT, ""
		case domain.BankCard:
			return pb.PaymentType_BANKCARD, ""
		}
		return pb.PaymentType_UNKNOW_PAYMENTTYPE, ""
	}
}

func convertStatus(isFromProto bool, pbStatus pb.Status, dbStatus domain.Status) (pb.Status, domain.Status) {
	if isFromProto {
		switch pbStatus {
		case pb.Status_TBC:
			return 0, domain.ToBeConfirmed
		case pb.Status_PROCESSING:
			return 0, domain.Processing
		case pb.Status_PAUSED:
			return 0, domain.Paused
		case pb.Status_SUCCEED:
			return 0, domain.Succeed
		case pb.Status_FAILED:
			return 0, domain.Failed
		}
		return 0, domain.UnknowStatus
	} else {
		switch dbStatus {
		case domain.ToBeConfirmed:
			return pb.Status_TBC, ""
		case domain.Processing:
			return pb.Status_PROCESSING, ""
		case domain.Paused:
			return pb.Status_PAUSED, ""
		case domain.Succeed:
			return pb.Status_SUCCEED, ""
		case domain.Failed:
			return pb.Status_FAILED, ""
		}
		return pb.Status_UNKNOW_STATUS, ""
	}
}

func convertMerchantRateType(isFromProto bool, pbMerchantRateType pb.MerchantRateType, dbMerchantRateType domain.MerchantRateType) (pb.MerchantRateType, domain.MerchantRateType) {
	if isFromProto {
		switch pbMerchantRateType {
		case pb.MerchantRateType_ratio:
			return 0, domain.Ratio
		case pb.MerchantRateType_fixed:
			return 0, domain.Fixed
		}
		return 0, domain.UnknowMerchantRateType
	} else {
		switch dbMerchantRateType {
		case domain.Ratio:
			return pb.MerchantRateType_ratio, ""
		case domain.Fixed:
			return pb.MerchantRateType_fixed, ""
		}
		return pb.MerchantRateType_unknow_merchant_rateType, ""
	}
}

func convertMerchantResult(isFromProto bool, pbMerchantResult pb.MerchantResult, dbMerchantResult domain.MerchantResultType) (pb.MerchantResult, domain.MerchantResultType) {
	if isFromProto {
		switch pbMerchantResult {
		case pb.MerchantResult_MerchantResultSucceed:
			return 0, domain.MerchantSucceed
		case pb.MerchantResult_MerchantResultFailed:
			return 0, domain.MerchantFailed
		}
		return 0, domain.UnknowMerchantResult
	} else {
		switch dbMerchantResult {
		case domain.MerchantSucceed:
			return pb.MerchantResult_MerchantResultSucceed, ""
		case domain.MerchantFailed:
			return pb.MerchantResult_MerchantResultFailed, ""
		}
		return pb.MerchantResult_UnknowMerchantResult, ""
	}
}

func timeConverter(t uint64) (*time.Time, error) {
	timeStr := strconv.FormatUint(t, 10)
	timeInt, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(timeInt, 0)
	return &tm, nil
}
