INSERT INTO `transaction` (`transaction_id`, `customer_name`, `transaction_dat`)
    VALUES("Trx001","Rivaldo","2022-10-22")

INSERT INTO `transaction` (`transaction_id`, `customer_name`, `transaction_dat`)
    VALUES("Trx002","Melina","2023-0407") 

INSERT INTO `detail_transaction` (`transaction_id`, `product_name`, `price`,`qty`,`discount`) 
VALUES ("Trx001", "Pan Pizza", "30.000",1 ,0.05),
       ( "Trx001", "Crowncrust", "45000",1 ,0.05),
       ( "Trx002", "Cheesy Bites", "47000",2, 0.1);



UPDATE `detail_transaction` SET discount = 2.5 WHERE transaction_id = "Trx001" AND product_name = "Pan Pizza";
UPDATE `detail_transaction` SET discount = 7.5 WHERE transaction_id = "Trx001" AND product_name = "Crowncrust";
UPDATE `detail_transaction` SET discount = 15 WHERE transaction_id = "Trx002" AND product_name= "Cheesy Bites";


SELECT id, transaction.customer_name, transaction.date,
       SUM(detail_transaction.qty) as qty, SUM(detail_transaction.price) as price, SUM(detail_transaction.discount*detail_transaction.price*detail_transaction.qty) as discount,
       SUM(detail_transaction.price * detail_transaction.qty - detail_transaction.discount*detail_transaction.price*detail_transaction.qty) as totalPrice
FROM transaction
JOIN detail_transaction ON transaction.transaction_id = detail_transaction.transaction_id
GROUP BY transaction.transaction_id;