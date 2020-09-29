-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `account_daily_counts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `daily_key` varchar(255) NOT NULL,
    `type` varchar(255) NOT NULL,
    `account` varchar(255) NOT NULL,
    `daily_count` decimal(20, 2) NOT NULL DEFAULT 0.00,
    `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uix_daily_key` (`daily_key`),
    INDEX `idx_account`(`account`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

ALTER TABLE `player_accounts` ADD COLUMN `merchant_id` int(10) unsigned NOT NULL;
ALTER TABLE `player_accounts` ADD INDEX `idx_merchant_id`(`merchant_id`) USING BTREE;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `account_daily_counts`;
ALTER TABLE `player_accounts` DROP COLUMN `merchant_id`;
DROP INDEX `idx_merchant_id` ON `player_accounts`;