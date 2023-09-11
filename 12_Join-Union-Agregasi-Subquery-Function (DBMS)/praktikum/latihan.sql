1.1
-- Insert 5 operators pada table operators. 
INSERT INTO operators (name, created_at, updated_at)
VALUES
    ('Operator 1', NOW(), NOW()),
    ('Operator 2', NOW(), NOW()),
    ('Operator 3', NOW(), NOW()),
    ('Operator 4', NOW(), NOW()),
    ('Operator 5', NOW(), NOW());

    -- Insert 3 product type.
INSERT INTO product_types (name,created_at,updated_at)VALUES
	('pakaian 1', NOW(), NOW()),
    ('pakaian 2', NOW(), NOW()),
    ('pakaian 3', NOW(), NOW());

-- Insert 2 product dengan product type id = 1, dan operators id = 3.
  INSERT INTO product (code, name, status, product_type_id, operator_id, created_at, updated_at)VALUES
    ('P001', 'Produk 1', 12, 1, 3,  NOW(), NOW()),
    ('P002', 'Produk 2', 5, 1, 3,  NOW(), NOW());

-- Insert 3 product dengan product type id = 2, dan operators id = 1.
  INSERT INTO product (code, name, status, product_type_id, operator_id, created_at, updated_at)VALUES
    ('P003', 'Produk 3', 2, 2, 1,  NOW(), NOW()),
    ('P004', 'Produk 4', 7, 2, 1,  NOW(), NOW()),
    ('P005', 'Produk 5', 3, 2, 1,  NOW(), NOW());
-- Insert 3 product dengan product type id = 3, dan operators id = 4.
  INSERT INTO product (code, name, status, product_type_id, operator_id, created_at, updated_at) VALUES
    ('P006', 'Produk 6', 6, 3, 4,  NOW(), NOW()),
    ('P007', 'Produk 7', 8, 3, 4,  NOW(), NOW()),
    ('P008', 'Produk 8', 0, 3, 4,  NOW(), NOW());

-- Insert product description pada setiap product.
  INSERT INTO product_description(deskripsi_produk,created_at, updated_at)VALUES
  ('ready Produk 1', NOW(), NOW()),
  ('sold out Produk 2', NOW(), NOW()),
  ('ready Produk 3', NOW(), NOW()),
  ('sold out Produk 4', NOW(), NOW()),
  ('ready Produk 5', NOW(), NOW()),
  ('sold out Produk 6', NOW(), NOW()),
  ('ready Produk 7', NOW(), NOW()),
  ('sold out Produk 8', NOW(), NOW());
  -- Insert 3 payment methods.
INSERT INTO payment_method(name, status, created_at, updated_at)
VALUES
    ('Metode Pembayaran 1', 3, NOW(), NOW()),
    ('Metode Pembayaran 2', 4, NOW(), NOW()),
    ('Metode Pembayaran 3', 5, NOW(), NOW()),
    ('Metode Pembayaran 4', 3, NOW(), NOW()),
    ('Metode Pembayaran 5', 7, NOW(), NOW()),
    ('Metode Pembayaran 6', 8, NOW(), NOW()),
    ('Metode Pembayaran 7', 2, NOW(), NOW()),
    ('Metode Pembayaran 8', 7, NOW(), NOW());
-- Insert 5 user pada tabel user.
INSERT INTO user (status, dob, gender, created_at, updated_at) VALUES
    (12, '1990-01-15', 'M', NOW(), NOW()),
    (11, '1985-05-20', 'F', NOW(), NOW()),
    (10, '1995-11-10', 'M', NOW(), NOW()),
    (6, '1982-03-25', 'F', NOW(), NOW()),
    (5, '1998-07-05', 'M', NOW(), NOW());

-- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

-- Pengguna 1
INSERT INTO transaction (user_id, status, total_qyt,total_price,payment_method_id,created_at, updated_at)
VALUES
    (1, 'Success', 3, 100.00, 1, NOW(), NOW()),
    (1, 'Pending', 2, 75.50, 2, NOW(), NOW()),
    (1, 'Failed', 1, 30.25, 3, NOW(), NOW());
-- Pengguna 2
INSERT INTO transaction (user_id, status, total_qyt,total_price,payment_method_id,created_at, updated_at)
VALUES
    (2,'Success', 4, 150.00, 4, NOW(), NOW()),
    (2,'Success', 1, 25.00, 5,NOW(), NOW()),
    (2,'Pending', 2, 75.50, 6,NOW(), NOW());
    -- Pengguna 3
INSERT INTO transaction (user_id, status, total_qyt, total_price, payment_method_id, created_at, updated_at)
VALUES
    (3, 'Success', 2, 150.00, 7, NOW(), NOW()),
    (3, 'Success', 4, 25.00, 8, NOW(), NOW()),
    (3, 'Pending', 3, 75.50, 9, NOW(), NOW());
-- Insert 3 product di masing-masing transaksi.

-- transaksi 1
INSERT INTO transaction_detail (transaction_detail_id,status, qty,price,product_id,created_at,updated_at) VALUES
    (1,'ready', 3, 50.00, 1,  NOW(), NOW()), 
    (1,'ready', 4, 20.00, 2, NOW(), NOW()), 
    (1,'ready', 3, 10.00, 3, NOW(), NOW()); 
    
    -- transaksi 2
INSERT INTO transaction_detail (transaction_detail_id,status, qty,price,product_id,created_at,updated_at) VALUES
    (2,'sold out', 2, 75.00, 4,  NOW(), NOW()), 
    (2,'sold out', 2, 25.00, 5, NOW(), NOW()), 
    (2,'sold out', 1, 150.00, 6, NOW(), NOW()); 

   -- transaksi 3
INSERT INTO transaction_detail (transaction_detail_id,status, qty,price,product_id,created_at,updated_at) VALUES
    (3,'ready', 2, 100.00, 7,  NOW(), NOW()), 
    (3,'sold out', 2, 50.00, 8, NOW(), NOW()), 
    (3,'ready', 3, 10.00, 9, NOW(), NOW()); 







-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT user_id, status
FROM user
WHERE gender = 'M' OR gender = 'Laki-laki';

-- Tampilkan product dengan id = 3.
 select * from product
 where product_id = 3;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM user
WHERE created_at >= DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY)
AND created_at <= CURRENT_DATE()
AND nama LIKE '%a%';


-- Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*)
FROM user
WHERE gender = 'F';
-- Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM user
ORDER BY nama ASC;
-- Tampilkan 5 data pada data product
SELECT *FROM product
LIMIT 5;
-- Ubah data product id 1 dengan nama ‘product dummy’. 
UPDATE product
SET name = 'product dummy'
WHERE product_id = 1;




-- Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_detail
SET qty = 3
WHERE product_id = 1;

-- Delete data pada tabel product dengan id = 1.
DELETE FROM product
WHERE product_id = 1;
-- Delete pada pada tabel product dengan product type id 1.
DELETE FROM product
WHERE product_type_id = 1;

2.2

-- Gabungkan data transaksi dari user id 1 dan user id 2.

-- transaksi dari user ID 1
SELECT * FROM onlineshop.transaction
WHERE user_id = 1
-- transaksi dari user ID 2
      UNION ALL
SELECT * FROM onlineshop.transaction
WHERE user_id = 2;

    -- Tampilkan jumlah harga transaksi user id 1.
SELECT user_id, SUM(total_price) AS transaction
FROM transaction
WHERE user_id = 1
GROUP BY user_id;

-- Tampilkan total transaksi dengan product type 2.
SELECT COUNT(*) AS total_transaksi_product_type2
FROM transaction_detail AS t
JOIN product AS p ON t.product_id = p.product_id
WHERE p.product_type_id = 2;


-- Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT *, pt.name AS product_type_name
FROM product AS p
JOIN product_types AS pt ON p.product_type_id = pt.product_type_id;

-- Tampilkan semua field table transaction, field name table product dan field name table user.

SELECT t.*, p.name AS product_name, u.nama AS user_name, td.product_id
FROM `transaction` AS t
JOIN user AS u ON t.user_id = t.user_id
JOIN transaction_detail AS td ON t.transaction_id = td.transaction_detail_id
JOIN product AS p ON td.product_id = td.product_id;


-- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER $$

CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transaction
FOR EACH ROW
BEGIN
    DELETE FROM transaction_detail
    WHERE transaction_id = OLD.transaction_id;
END;
$$
DELIMITER ;


-- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

DELIMITER $$

CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transaction_detail
FOR EACH ROW
BEGIN
    DECLARE old_transaction_id INT;
    DECLARE old_qty INT;
    
    -- Simpan nilai kolom yang dihapus
    SET old_transaction_id = OLD.transaction_id;
    SET old_qty = OLD.qty;
    
    -- Perbarui total_qty dalam tabel yang sesuai
    UPDATE total_qty_table
    SET total_qty = total_qty - old_qty
    WHERE transaction_id = old_transaction_id;
END;
$$

DELIMITER ;


-- tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. 

SELECT * FROM product
WHERE product_id NOT IN (SELECT DISTINCT product_id FROM transaction_detail);