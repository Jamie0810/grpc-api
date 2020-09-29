package domain

import "time"

type OrderHistory struct {
	ID              int64      `gorm:"column:id;type:serial;PRIMARY_KEY;AUTO_INCREMENT:true;NOT NULL"`
	TrackingNumber  string     `gorm:"column:tracking_number;type:varchar(20);INDEX:idx_tracking_number;NOT NULL"` // 訂單編號
	ChannelName     string     `gorm:"column:channel_name;type:varchar(30);NOT NULL"`                              // 渠道名稱
	IsSuccess       bool       `gorm:"column:is_success;type:boolean;NOT NULL;DEFAULT:false"`                      // 建立結果 e.g 成功 / 失敗
	ResponseMessage string     `gorm:"column:response_message;type:LONGTEXT;NOT NULL;DEFAULT:''"`                  // 錯誤訊息
	CreatedAt       time.Time  `gorm:"column:created_at;type:datetime;NOT NULL;DEFAULT:now()"`                     // 建立時間
	UpdatedAt       time.Time  `gorm:"column:updated_at;type:datetime;DEFAULT:NULL"`                               // 欄位更新時間
	DeletedAt       *time.Time `gorm:"column:deleted_at;type:datetime;DEFAULT:NULL"`                               // 軟刪除時間
}
