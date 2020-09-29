-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO `order_channels` (`id`, `channel_id`, `channel_name`, `is_enable`)
VALUES
    (8, 'vgam', 'VGAM', 1),
    (9, 'atp', 'VATP', 1);

INSERT INTO `order_channels` (`id`, `channel_id`, `channel_name`, `is_enable`)
VALUES
    (10, 'shan6', 'SHAN6', 1);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM `order_channels` WHERE `id` IN (8,9) ;
DELETE FROM `order_channels` WHERE `id` IN (10);
