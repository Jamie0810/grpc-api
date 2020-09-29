-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `sub_order_histories` ADD COLUMN `channel_req` json NULL AFTER `response_message`;
ALTER TABLE `sub_order_histories` ADD COLUMN `channel_resp` json NULL AFTER `channel_req`;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `sub_order_histories` DROP COLUMN `channel_req`;
ALTER TABLE `sub_order_histories` DROP COLUMN `channel_resp`;