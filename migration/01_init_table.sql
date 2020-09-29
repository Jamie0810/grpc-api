-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `orders` (
  `tracking_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `wallet_two_phase_status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'before',
  `transacion_type` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `payment_number` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '-',
  `payment_type` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `deposit_account` json NULL,
  `withdraw_account` json NULL,
  `status` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'processing',
  `player_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '-',
  `estimated_cost` decimal(20, 2) NOT NULL DEFAULT 0.00,
  `actual_cost` decimal(20, 2) NOT NULL DEFAULT 0.00,
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `completed_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `merchant_id` bigint(20) unsigned NOT NULL,
  `merchant_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `merchant_fee` decimal(20, 2) NOT NULL DEFAULT 0.00,
  `merchant_rate_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'fixed',
  `merchant_rate` decimal(20, 2) NOT NULL DEFAULT 0.00,
  `merchant_rate_fixed` decimal(20, 2) NOT NULL DEFAULT 0.00,
  `merchant_project_id` bigint(20) unsigned NOT NULL,
  `merchant_project_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `merchant_product_id` bigint(20) unsigned NOT NULL,
  `merchant_product_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `merchant_order_number` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `merchant_result` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '-',
  `merchant_notify_url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_order_number` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '-',
  `channel_merchant_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`tracking_number`) USING BTREE,
  INDEX `idx_order_created_at`(`created_at`) USING BTREE,
  INDEX `idx_order_completed_at`(`completed_at`) USING BTREE,
  INDEX `idx_order_updated_at`(`updated_at`) USING BTREE,
  INDEX `idx_merchant_order_number`(`merchant_order_number`) USING BTREE,
  INDEX `idx_merchant_project_id`(`merchant_project_id`) USING BTREE,
  INDEX `idx_merchant_product_id`(`merchant_product_id`) USING BTREE,
  INDEX `idx_channel_order_number`(`channel_order_number`) USING BTREE,
  INDEX `idx_channel_merchant_id` (`channel_merchant_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
CREATE TABLE `order_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tracking_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_success` tinyint(1) NOT NULL DEFAULT 0,
  `response_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  INDEX `idx_tracking_number`(`tracking_number`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
CREATE TABLE `sub_order_histories` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tracking_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sub_tracking_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_success` tinyint(1) NOT NULL DEFAULT 0,
  `response_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  INDEX `idx_tracking_number_sub_tracking_number`(`tracking_number`, `sub_tracking_number`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
CREATE TABLE `order_channels` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_enable` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
CREATE TABLE `mock_payment_infos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `host` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` int(10) NOT NULL DEFAULT 0,
  `message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `notify_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `error` tinyint(1) NOT NULL DEFAULT 0,
  `retry` int(10) unsigned NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
CREATE TABLE `polypay_params` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `merchant_project_id` bigint(20) unsigned NOT NULL,
  `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `host` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `polypay_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `polypay_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `origin` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `origin_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_enable` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  INDEX `idx_merchant_project_id`(`merchant_project_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
INSERT INTO `order_channels` (`id`, `channel_id`, `channel_name`, `is_enable`)
VALUES
	(1,'mock-polypay','聚合支付Mock',1),
	(2,'mock-shan3','閃3Mock',0),
	(3,'mock-shan6','閃6Mock',0),
	(4,'mock-bankcard','銀行卡Mock',0),
	(5,'mock-beiming1','北冥1Mock',0),
	(6,'mock-beiming2','北冥2Mock',0),
	(7,'mock-exchange','兌換支付',0);
INSERT INTO `mock_payment_infos` (`id`, `channel_id`, `channel_name`, `host`, `status`, `message`, `notify_url`, `error`, `retry`)
VALUES
	(1,'mock-polypay','聚合支付','https://mockserver-gbhzzlku2a-de.a.run.app/deposite/polypay?pay_url=https://t.cn/pay&account=13838383838&name=alex&trade_amount=0.01&receipt_amount=0.01&upstream_order=201800125481245844575112541&repeat_pay=false&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443',10000,'success','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-polypay',0,10),
	(2,'mock-bankcard','銀行卡','https://mockserver-gbhzzlku2a-de.a.run.app/withdraw/bankcard?flash_id=fla1234&merchant=A0001&payed_money=1000&merchant_order_id=ord1234&payed_time=2020-01-02 15:15:03&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443',10000,'请求成功','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-backcard',0,10),
	(3,'mock-shan3','閃3','https://mockserver-gbhzzlku2a-de.a.run.app/deposite/shan3?pay_url=http://127.0.0.1/api/pay.jsp?id=2%26c=e920dc57680ce4bf2b89654440adfee5&tid=546564&trade_amount=0.01&receipt_amount=0.01&user_data=abcd&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443',10000,'success','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-shan3',0,10),
	(4,'mock-shan6','閃6','https://mockserver-gbhzzlku2a-de.a.run.app/deposite/shan6?flash_id=DA20170929042702499412xxxxxxxx&qr_code=https://qr.alipay.com/xxxxxxxxxx&payee_name=alex&expire=2000-01-02 03:04:05&merchant=xxxx&payed_money=9.78&upstream_order=201800125481245844575112541&repeat_pay=no&merchant_order_id=368571021&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443',10000,'订单生成成功','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-shan6',0,10),
	(5,'mock-beiming1','北冥1','https://mockserver-gbhzzlku2a-de.a.run.app/deposite/beiming1?merchant_code=ONION&merchant_order=test_order3&flash_id=DA20170929042702499412xxxxxxxx&pay_url=https://receipt.trqsxt.com/receipt-html/wait/&bill_price=10&card_num=9560580597580076&accept_name=alex&bank_name=支付宝支付&province=&city=taipei&country=taiwan&node=taipei branch&created_at=2020-01-02 15:15:03&remark=remark&receipt_price=10&payed_time=2020-05-04 15:54:13&order_type=2&order_mode=0&user_level=2&target_account=9560580597580076&target_name=alex&trade_time=2020-01-02 15:15:03&repeat_pay&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443',1,'创建成功','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-beiming1',0,10),
	(6,'mock-beiming2','北冥2','https://mockserver-gbhzzlku2a-de.a.run.app/deposite/beiming2?merchant_code=ONION&merchant_order=test_order3&flash_id=123456&pay_url=https://receipt.trqsxt.com/receipt-html/wait/&bill_price=10&card_num=9560580597580076&accept_name=alex&bank_name=支付宝支付&province=&city=taipei&country=taiwan&node=taipei branch&created_at=2020-01-02 15:15:03&receipt_price=10&payed_time=2020-05-04 15:54:13&target_account=9560580597580076&target_name=alex&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443',1,'订单创建成功','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-beiming2',0,10),
	(7,'mock-exchange','兌換支付','https://mockserver-gbhzzlku2a-de.a.run.app/withdraw/exchange?body=Success&order_id=ord1234&trade_No=20170113110070001502510003065349&service_charge=0&order_amout=100&pay_account=lusu@qq.com&pay_time=2020-01-02 15:15:03&sign=1cf245d08e3a2fac8fc93ab964542a1&t=1492257443',10000,'Success','https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-exchange',0,10);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `mock_payment_infos`;
DROP TABLE IF EXISTS `orders`;
DROP TABLE IF EXISTS `order_histories`;
DROP TABLE IF EXISTS `sub_order_histories`;
DROP TABLE IF EXISTS `order_channels`;
DROP TABLE IF EXISTS `polypay_params`;

