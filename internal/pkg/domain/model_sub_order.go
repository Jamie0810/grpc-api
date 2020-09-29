package domain

import (
	"time"
)

type SubOrder struct {
	ID                     int64      `gorm:"column:id;type:serial;PRIMARY_KEY;AUTO_INCREMENT:true;NOT NULL"`
	TrackingNumber         string     `gorm:"column:tracking_number;type:varchar(20);INDEX:idx_tracking_number_sub_tracking_number;NOT NULL"`     // 訂單編號
	SubTrackingNumber      string     `gorm:"column:sub_tracking_number;type:varchar(20);INDEX:idx_tracking_number_sub_tracking_number;NOT NULL"` // 子訂單編號
	ChannelID              string     `gorm:"column:channel_id;type:varchar(40);NOT NULL"`                                                        // 渠道id
	ChannelName            string     `gorm:"column:channel_name;type:varchar(30);NOT NULL"`                                                      // 渠道名稱
	MerchantProductAPIName string     `gorm:"column:merchant_product_api_name;type:varchar(50);NOT NULL"`                                         // 商戶接口名稱
	IsSuccess              bool       `gorm:"column:is_success;type:boolean;NOT NULL;DEFAULT:false"`                                              // 建立結果 e.g 成功 / 失敗
	ResponseMessage        string     `gorm:"column:response_message;type:LONGTEXT;NOT NULL;DEFAULT:''"`                                          // 錯誤訊息
	ChannelReq             string     `gorm:"column:channel_req;type:LONGTEXT;DEFAULT:''"`                                                        // 渠道建單參數
	ChannelResp            string     `gorm:"column:channel_resp;type:LONGTEXT;DEFAULT:''"`                                                       // 渠道回應結果
	CreatedAt              time.Time  `gorm:"column:created_at;type:datetime;NOT NULL;DEFAULT:now()"`                                             // 建立時間
	UpdatedAt              time.Time  `gorm:"column:updated_at;type:datetime;DEFAULT:NULL"`                                                       // 欄位更新時間
	DeletedAt              *time.Time `gorm:"column:deleted_at;type:datetime;DEFAULT:NULL"`                                                       // 軟刪除時間
}
