-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `player_accounts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `player_id` varchar(30)  NOT NULL,
    `deposit_account` varchar(255) NULL,
    `withdraw_account` varchar(255) NULL,
    `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_player_id`(`player_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


CREATE TABLE `player_daily_counts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `daily_key` varchar(255) NOT NULL,
    `player_id` varchar(30)  NOT NULL,
    `merchant_id` int(10) unsigned NOT NULL,
    `deposit_daily_count` decimal(20, 2) NOT NULL DEFAULT 0.00,
    `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uix_daily_key` (`daily_key`),
    INDEX `idx_player_id`(`player_id`) USING BTREE,
    INDEX `idx_merchant_id`(`merchant_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

CREATE TABLE `merchant_daily_counts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `daily_key` varchar(255) NOT NULL,
    `merchant_id` int(10) unsigned NOT NULL,
    `deposit_daily_count` decimal(20, 2) NOT NULL DEFAULT 0.00,
    `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uix_daily_key` (`daily_key`),
    INDEX `idx_merchant_id`(`merchant_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `player_accounts`;
DROP TABLE IF EXISTS `player_daily_counts`;
DROP TABLE IF EXISTS `merchant_daily_counts`;