-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` ADD COLUMN `expired_at` datetime(0) NULL AFTER `deleted_at`;
ALTER TABLE `orders` ADD COLUMN `channel_req` json NULL AFTER `channel_name`;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `orders` DROP COLUMN `expired_at`;
ALTER TABLE `orders` DROP COLUMN `channel_req`;