-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `wallets` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL,
  `merchant_id` int(10) unsigned NOT NULL,
  `project_id` int(10) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `balance` decimal(20,2) NOT NULL DEFAULT '0.00',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_wallets_key` (`key`),
  KEY `idx_wallets_merchant_id` (`merchant_id`),
  KEY `idx_wallets_project_id` (`project_id`),
  KEY `idx_wallets_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `wallets` (`id`, `key`, `merchant_id`, `project_id`, `product_id`, `balance`, `created_at`, `updated_at`)
VALUES
	(1,'bs40isrruu7mf8t09sn0',1,1,1,19000000.00,'2020-07-10 06:23:47','2020-07-15 08:19:52');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS `wallets`;