-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `wallet_histories` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL,
  `merchant_id` int(10) unsigned NOT NULL,
  `project_id` int(10) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `wallet_event` varchar(50) NOT NULL,
  `fee_rate_type` varchar(50) NOT NULL,
  `fee_rate` decimal(20,2) NOT NULL,
  `source_type` varchar(50) NOT NULL,
  `source_table` varchar(50) NOT NULL,
  `source_table_key` varchar(50) NOT NULL,
  `source_table_offset` varchar(50) NOT NULL,
  `origin_balance` decimal(20,2) NOT NULL DEFAULT '0.00',
  `operation_balance` decimal(20,2) NOT NULL DEFAULT '0.00',
  `final_balance` decimal(20,2) NOT NULL DEFAULT '0.00',
  `operator_name` varchar(255) DEFAULT NULL,
  `operator_key` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_wallet_histories_key` (`key`),
  UNIQUE KEY `uix_wallet_histories_source_event` (`source_table_key`,`source_table_offset`,`wallet_event`,`source_table`),
  KEY `idx_wallet_histories_merchant_id` (`merchant_id`),
  KEY `idx_wallet_histories_project_id` (`project_id`),
  KEY `idx_wallet_histories_product_id` (`product_id`),
  KEY `idx_wallet_histories_wallet_event` (`wallet_event`),
  KEY `idx_wallet_histories_source_type` (`source_type`),
  KEY `idx_wallet_histories_operator_name` (`operator_name`),
  KEY `idx_wallet_histories_operator_key` (`operator_key`),
  KEY `idx_wallet_histories_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `wallet_histories`;