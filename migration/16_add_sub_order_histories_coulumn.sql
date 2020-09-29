-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `sub_order_histories` ADD COLUMN `merchant_product_api_name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL AFTER `channel_name`;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `sub_order_histories` DROP COLUMN `merchant_product_api_name`;
