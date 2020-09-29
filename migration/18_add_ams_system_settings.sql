-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `system_settings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `maximum_fee_rate` decimal(5,2) NOT NULL DEFAULT '100.00',
  `minimum_fee_rate` decimal(5,2) NOT NULL DEFAULT '0.00',
  `maximum_fee` decimal(20,2) NOT NULL DEFAULT '9999999.99',
  `minimum_fee` decimal(20,2) NOT NULL DEFAULT '0.00',
  `maximum_withdraw` decimal(20,2) NOT NULL DEFAULT '9999999.99',
  `minimum_withdraw` decimal(20,2) NOT NULL DEFAULT '0.00',
  `security_deposit` decimal(20,2) NOT NULL DEFAULT '0.00',
  `retry_on_order_failed` decimal(20,0) NOT NULL DEFAULT '0',
  `case_closed_time` int(11) NOT NULL DEFAULT '24',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `system_settings` (`id`, `maximum_fee_rate`, `minimum_fee_rate`, `maximum_fee`, `minimum_fee`, `maximum_withdraw`, `minimum_withdraw`, `security_deposit`, `retry_on_order_failed`, `case_closed_time`)
VALUES
	(1,2.00,1.00,2.00,1.00,999999.99,1.00,301.00,5,1),
	(2,100.00,0.00,9999999.99,0.00,9999999.99,0.00,0.00,0,24);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `system_settings`;