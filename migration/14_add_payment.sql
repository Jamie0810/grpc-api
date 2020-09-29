-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `payment_params` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `merchant_product_id` bigint(20) unsigned NOT NULL,
  `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `host` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `payment_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `payment_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `origin` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `origin_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `is_enable` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  INDEX `idx_merchant_product_id`(`merchant_product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
INSERT INTO `payment_params` (`id`, `merchant_product_id`, `channel_id`, `channel_name`, `host`, `payment_id`, `payment_token`, `origin`, `origin_token`, `is_enable`)
VALUES
	(1,1,'polypay','聚合支付','http://203.95.192.98:52003/api/v1/trade/mock-polypay-create','MOCK-ID-5b03835a-18d7-47da-896b-808bcf1e6de5','MOCK-TOKEN-0da1f44b-5253-477a-a9ef-fa91e3488ff8','origin','token',1),
	(2,1,'bankcard','銀行卡','http://10.205.32.23','','','TWGATEWAY','',1),
	(3,1,'beiming1','北冥1','http://10.205.60.37:54555','','','TWGATEWAY','',1),
	(4,1,'beiming2','北冥2','http://10.205.60.35:54555','','','TWGATEWAT','',1),
	(5,1,'exchange','兌換支付','https://trans2.iwyfu.com/m/','','','TWGATEWAY','',1),
	(6,1,'vgam','vGAM','https://sit-vgam.gcp.silkrode.com.tw','14','e016c9e3-ec57-40ef-baba-1e6d9e50ab58','','',1),
	(7,1,'atp','ATP','https://sit-vatp-api.silkrode.in','131','dc6c7982-4442-4010-aa1d-64c01bb1f44b','','',1);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `payment_params`;
