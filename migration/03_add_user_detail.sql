-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` ADD COLUMN `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' AFTER `player_id`;
ALTER TABLE `orders` ADD COLUMN `device_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' AFTER `user_name`;
ALTER TABLE `orders` ADD COLUMN `device_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' AFTER `device_id`;
ALTER TABLE `orders` ADD COLUMN `device_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' AFTER `device_ip`;
ALTER TABLE `orders` ADD COLUMN `telephone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' AFTER `device_type`;
ALTER TABLE `orders` ADD COLUMN `merchant_product_api_id` bigint(20) unsigned NOT NULL AFTER `merchant_product_name`;
ALTER TABLE `orders` ADD COLUMN `merchant_product_api_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL AFTER `merchant_product_api_id`;
ALTER TABLE `orders` ADD INDEX `idx_merchant_product_api_id`(`merchant_product_api_id`) USING BTREE;
ALTER TABLE `orders` ADD INDEX `idx_merchant_product_api_name`(`merchant_product_api_name`) USING BTREE;
ALTER TABLE `orders` ADD INDEX `idx_merchant_product_name`(`merchant_product_name`) USING BTREE;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX `idx_merchant_product_api_id`  ON `orders`;
DROP INDEX `idx_merchant_product_api_name` ON `orders`;
ALTER TABLE `orders` DROP COLUMN `user_name`;
ALTER TABLE `orders` DROP COLUMN `device_id`;
ALTER TABLE `orders` DROP COLUMN `device_ip`;
ALTER TABLE `orders` DROP COLUMN `device_type`;
ALTER TABLE `orders` DROP COLUMN `telephone`;
ALTER TABLE `orders` DROP COLUMN `merchant_product_api_id`;
ALTER TABLE `orders` DROP COLUMN `merchant_product_api_name`;