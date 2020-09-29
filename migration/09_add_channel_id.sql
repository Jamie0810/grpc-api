-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` CHANGE `channel_id` `channel_code` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL;
ALTER TABLE `orders` ADD COLUMN `channel_id` bigint(20) unsigned NOT NULL AFTER `merchant_notify_url`;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `orders` DROP COLUMN `channel_id`;
ALTER TABLE `orders` CHANGE `channel_code` `channel_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL;