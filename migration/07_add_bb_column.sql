-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `orders` ADD COLUMN `bb_reason` json NULL;
ALTER TABLE `orders` ADD COLUMN `bb_skip` json NULL;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `orders` DROP COLUMN `bb_reason`;
ALTER TABLE `orders` DROP COLUMN `bb_skip`;