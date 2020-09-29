-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` CHANGE `status` `status` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'PROCESSING';
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `orders` CHANGE `status` `status` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'PROCESSING';