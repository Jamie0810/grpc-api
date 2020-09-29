-- +goose Up
-- SQL in this section is executed when the migration is applied.
UPDATE `mock_payment_infos` 
    SET `host` = 'https://mockserver-gbhzzlku2a-de.a.run.app/deposite/polypay?pay_url=https://t.cn/pay&account=13838383838&name=alex&trade_amount=0.5&receipt_amount=0.33&upstream_order=201800125481245844575112541&repeat_pay=false&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443&order_fee=10' 
    WHERE `id` = 1;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
UPDATE `mock_payment_infos` 
    SET `host` = 'https://mockserver-gbhzzlku2a-de.a.run.app/deposite/polypay?pay_url=https://t.cn/pay&account=13838383838&name=alex&trade_amount=0.5&receipt_amount=0.33&upstream_order=201800125481245844575112541&repeat_pay=false&sign=1cf245d08e3a2fac8fc93ab964542a1&time=1492257443&order_fee=10' 
    WHERE `id` = 1;