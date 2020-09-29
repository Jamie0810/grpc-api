-- +goose Up
-- SQL in this section is executed when the migration is applied.
UPDATE `order_channels` SET `is_enable` = 1 WHERE `id` = 4;
UPDATE `mock_payment_infos` 
    SET `host`='https://mockserver-gbhzzlku2a-de.a.run.app/withdraw/bankcard?merchant=A0001&payed_money=1001&payed_time=2020-01-02%2015:15:03&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443', `notify_url`='https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-bankcard'
    WHERE `id` = 2;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
UPDATE `order_channels` SET `is_enable` = 1 WHERE `id` = 4;
UPDATE `mock_payment_infos` 
    SET `host`='https://mockserver-gbhzzlku2a-de.a.run.app/withdraw/bankcard?merchant=A0001&payed_money=1001&payed_time=2020-01-02%2015:15:03&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443', `notify_url`='https://api.huimai027.com/public/api/v1/order/thridPartyNotification/mock-bankcard'
    WHERE `id` = 2;