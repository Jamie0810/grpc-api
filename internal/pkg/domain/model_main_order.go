package domain

// These models are only used to store into database
import (
	"encoding/json"
	"time"
)

type MainOrder struct {
	TrackingNumber  string          `json:"TrackingNumber" gorm:"column:tracking_number;type:varchar(20);PRIMARY_KEY;NOT NULL"`                  // 訂單編號 (xid)
	WalletStatus    WalletStatus    `json:"WalletStatus" gorm:"column:wallet_two_phase_status;type:varchar(20);NOT NULL;DEFAULT:'before'"`       // 預扣錢包狀態
	TransactionType TransactionType `json:"TransactionType" gorm:"column:transacion_type;type:varchar(15);NOT NULL"`                             // 交易方式
	PaymentNumber   string          `json:"PaymentNumber" gorm:"column:payment_number;type:varchar(30);NOT NULL;DEFAULT:'-'"`                    // 支付流水號
	PaymentType     PaymentType     `json:"PaymentType" gorm:"column:payment_type;type:varchar(15);NOT NULL"`                                    // 訂單類型 e.g.微信 wechat / 支付寶 alipay / 銀行卡 back_card
	DepositAccount  string          `json:"DepositAccount" gorm:"column:deposit_account;type:LONGTEXT;DEFAULT:''"`                               // 打款帳戶
	WithdrawAccount string          `json:"WithdrawAccount" gorm:"column:withdraw_account;type:LONGTEXT;DEFAULT:''"`                             // 收款帳戶
	Status          Status          `json:"Status" gorm:"column:status;type:varchar(15);NOT NULL;DEFAULT:'processing'"`                          // 訂單狀態
	PlayerID        string          `json:"PlayerID" gorm:"column:player_id;type:varchar(30);NOT NULL;DEFAULT:'-'"`                              // 玩家id  KM
	UserName        string          `json:"UserName" gorm:"column:user_name;type:varchar(50);NOT NULL;DEFAULT:''"`                               // 玩家姓名 記錄用
	DeviceID        string          `json:"DeviceID" gorm:"column:device_id;type:varchar(50);NOT NULL;DEFAULT:''"`                               // 設備id 記錄用
	DeviceIP        string          `json:"DeviceIP" gorm:"column:device_ip;type:varchar(50);NOT NULL;DEFAULT:''"`                               // 設備ip 記錄用
	DeviceType      string          `json:"DeviceType" gorm:"column:device_type;type:varchar(20);NOT NULL;DEFAULT:''"`                           // 設備類型 記錄用
	Telephone       string          `json:"Telephone" gorm:"column:telephone;type:varchar(20);NOT NULL;DEFAULT:''"`                              // 電話號碼 記錄用
	EstimatedCost   float64         `json:"EstimatedCost" gorm:"column:estimated_cost;type:decimal(20,2);NOT NULL;DEFAULT:0.0"`                  // 預計交易金額
	ActualCost      float64         `json:"ActualCost" gorm:"column:actual_cost;type:decimal(20,2);NOT NULL;DEFAULT:0.0"`                        // 實際交易金額
	CreatedAt       time.Time       `json:"CreatedAt" gorm:"column:created_at;tyape:datetime;INDEX:idx_order_created_at;NOT NULL;DEFAULT:now()"` // 訂單建立時間
	CompletedAt     time.Time       `json:"CompletedAt" gorm:"column:completed_at;type:datetime;INDEX:idx_order_completed_at;DEFAULT:NULL"`      // 訂單完成時間
	UpdatedAt       time.Time       `json:"UpdatedAt" gorm:"column:updated_at;type:datetime;INDEX:idx_order_updated_at;DEFAULT:NULL"`            // 訂單更新時間
	DeletedAt       *time.Time      `json:"DeletedAt" gorm:"column:deleted_at;type:datetime;DEFAULT:NULL"`                                       // 軟刪除時間
	ExpiredAt       *time.Time      `json:"ExpiredAt" gorm:"column:expired_at;type:datetime;DEFAULT:NULL"`                                       // 訂單過期時間
	// customer 商戶
	MerchantID          uint64             `json:"MerchantID" gorm:"column:merchant_id;type:bigint;INDEX:idx_merchant_id;NOT NULL"`                                   // 商戶流水號 前端自組
	MerchantName        string             `json:"MerchantName" gorm:"column:merchant_name;type:varchar(30);NOT NULL"`                                                // 商戶姓名
	MerchantOrderNumber string             `json:"MerchantOrderNumber" gorm:"column:merchant_order_number;type:varchar(30);INDEX:idx_merchant_order_number;NOT NULL"` // 商戶訂單編號
	MerchantFee         float64            `json:"MerchantFee" gorm:"column:merchant_fee;type:decimal(20,2);NOT NULL;DEFAULT:0.0"`                                    // 商戶手續費
	MerchantRateType    MerchantRateType   `json:"MerchantRateType" gorm:"column:merchant_rate_type;type:varchar(10);NOT NULL;DEFAULT:'fixed'"`                       // 商戶費率選項 0: 未設定 1: 比率 2: 固定費率
	MerchantRate        float64            `json:"MerchantRate" gorm:"column:merchant_rate;type:decimal(20,2);NOT NULL;DEFAULT:0.0"`                                  // 商戶費率比率  KM
	MerchantRateFixed   float64            `json:"MerchantRateFixed" gorm:"column:merchant_rate_fixed;type:decimal(20,2);NOT NULL;DEFAULT:0.0"`                       // 商戶費率固定  KM
	MerchantResult      MerchantResultType `json:"MerchantResult" gorm:"column:merchant_result;type:varchar(30);NOT NULL;DEFAULT:'-'"`                                // 商戶回調結果
	MerchantNotifyURL   string             `json:"MerchantNotifyURL" gorm:"column:merchant_notify_url;type:varchar(100);NOT NULL"`                                    // 商戶回調網址
	// 商戶項目
	MerchantProjectID   uint64 `json:"MerchantProjectID" gorm:"column:merchant_project_id;type:bigint;INDEX:idx_merchant_project_id;NOT NULL"` // 商戶項目流水號
	MerchantProjectName string `json:"MerchantProjectName" gorm:"column:merchant_project_name;type:varchar(30);NOT NULL"`                      // 商戶項目名稱
	// 商戶產品
	MerchantProductID   uint64 `json:"MerchantProductID" gorm:"column:merchant_product_id;type:bigint;NOT NULL"`          // 產品id
	MerchantProductName string `json:"MerchantProductName" gorm:"column:merchant_product_name;type:varchar(30);NOT NULL"` // 產品名稱
	// 商戶接口
	MerchantProductAPIID   uint64 `json:"MerchantProductAPIID" gorm:"column:merchant_product_api_id;type:bigint;NOT NULL"`          // 商戶接口id
	MerchantProductAPIName string `json:"MerchantProductAPIName" gorm:"column:merchant_product_api_name;type:varchar(50);NOT NULL"` // 商戶接口名稱
	// 3rd party 支付渠道
	ChannelID          uint64          `json:"ChannelID" gorm:"column:channel_id;type:bigint;NOT NULL"`          // 渠道id
	ChannelCode        string          `json:"ChannelCode" gorm:"column:channel_code;type:varchar(40);NOT NULL"` // 渠道code
	ChannelName        string          `json:"ChannelName" gorm:"column:channel_name;type:varchar(30);NOT NULL"`
	ChannelReq         json.RawMessage `json:"ChannelReq" gorm:"column:channel_req;type:json"`                                                                             // 渠道建單參數
	ChannelOrderNumber string          `json:"ChannelOrderNumber" gorm:"column:channel_order_number;type:varchar(30);INDEX:idx_channel_order_number;NOT NULL;DEFAULT:'-'"` // 渠道訂單編號
	ChannelMerchantID  string          `json:"ChannelMerchantID" gorm:"column:channel_merchant_id;type:varchar(100);INDEX:idx_channel_merchant_id;NOT NULL"`               // 商戶原始
	// 風控暫停
	BBReason json.RawMessage `json:"BBReason" gorm:"column:bb_reason;type:json"` // 風控未通過原因
	BBSkip   json.RawMessage `json:"BBSkip" gorm:"column:bb_skip;type:json"`     // 風控要跳過步驟
}

type WalletStatus string
type TransactionType string
type PaymentType string
type Status string
type MerchantRateType string
type MerchantResultType string

// 預扣錢包狀態
const (
	UnknowWalletStatus WalletStatus = "unknow"
	BeforeNotify       WalletStatus = "before"
	ConfirmNotify      WalletStatus = "confirm"
	FailedNotify       WalletStatus = "failed"
)

// 訂單交易方式
const (
	// Deposit 商戶打款
	Deposit TransactionType = "DEPOSIT"
	// Withdrawal 商戶收款
	Withdrawal TransactionType = "WITHDRAW"
)

// 訂單類型
const (
	UnknowPaymentType PaymentType = "UNKNOW"
	// AliPay 支付寶
	AliPay PaymentType = "ALIPAY"
	// Wechat 微信
	Wechat PaymentType = "WECHAT"
	// BankCard 銀行卡
	BankCard PaymentType = "BANKCARD"
)

// 訂單狀態
const (
	UnknowStatus = "UNKNOW"
	// ToBeConfirmed 待確認
	ToBeConfirmed Status = "TBC"
	// Processing 處理中
	Processing Status = "PROCESSING"
	// Paused 風控暫停
	Paused Status = "PAUSED"
	// Succeed 交易成功
	Succeed Status = "SUCCEED"
	// Failed 交易失敗
	Failed Status = "FAILED"
)

// 手續費狀態
const (
	// 訂單費率比例
	UnknowMerchantRateType MerchantRateType = "unknow"
	Ratio                  MerchantRateType = "ratio"
	// 訂單費率固定
	Fixed MerchantRateType = "fixed"
)

// 商戶回調狀態
const (
	UnknowMerchantResult MerchantResultType = "UNKNOW"
	// Succeed 交易成功
	MerchantSucceed MerchantResultType = "SUCCEED"
	// Failed 交易失敗
	MerchantFailed MerchantResultType = "FAILED"
)
