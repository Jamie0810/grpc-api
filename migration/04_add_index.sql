-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` ADD INDEX `idx_status`(`status`) USING BTREE;
ALTER TABLE `orders` ADD INDEX `idx_merchant_id_player_id_created_at` (`merchant_id` , `player_id` , `created_at`) USING BTREE;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX `idx_status` ON `orders`;
DROP INDEX `idx_merchant_id_player_id_created_at` ON `orders`;