-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` CHANGE `deposit_account` `deposit_account` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE `orders` CHANGE `withdraw_account` `withdraw_account` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

ALTER TABLE `sub_order_histories` CHANGE `channel_req` `channel_req` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
ALTER TABLE `sub_order_histories` CHANGE `channel_resp` `channel_resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `orders` CHANGE `deposit_account` `deposit_account` json DEFAULT NULL;
ALTER TABLE `orders` CHANGE `withdraw_account` `withdraw_account` json DEFAULT NULL;


ALTER TABLE `sub_order_histories` CHANGE `channel_req` `channel_req` json DEFAULT NULL;
ALTER TABLE `sub_order_histories` CHANGE `channel_resp` `channel_resp` json DEFAULT NULL;