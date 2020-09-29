-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `main_order_settings` (
    `tracking_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `last_sub_tracking_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `risk_resp` json NULL,
    `cursor` tinyint(1) NOT NULL DEFAULT 0,
    `retry_times` tinyint(1) NOT NULL DEFAULT 0,
    `is_ended` tinyint(1) NOT NULL DEFAULT 0,
    `created_at` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updated_at` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
    `deleted_at` datetime(0) NULL,
    PRIMARY KEY (`tracking_number`) USING BTREE,
    INDEX `idx_last_sub_tracking_number`(`last_sub_tracking_number`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `main_order_settings`;