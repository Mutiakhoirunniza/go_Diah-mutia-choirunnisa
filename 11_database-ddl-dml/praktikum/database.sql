-- jawaban no 1 
CREATE DATABASE alta_online_shop;
use alta_online_shop;
-- jawaban no 2 a
CREATE TABLE user (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(120) NOT NULL,
    nama_lengkap VARCHAR(50),
    gender ENUM('pria', 'wanita'),
    alamat VARCHAR(255),
    tanggal_registrasi TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (username),
    UNIQUE (email)
);
			-- jawaban no 2 b 
            
   -- Tabel untuk tipe produk
CREATE TABLE IF NOT EXISTS product_type (
    product_type_id INT AUTO_INCREMENT PRIMARY KEY,
    nama_tipe VARCHAR(50) NOT NULL
);

-- Tabel untuk operator
CREATE TABLE IF NOT EXISTS operators (
    operator_id INT AUTO_INCREMENT PRIMARY KEY,
    nama_operator VARCHAR(50) NOT NULL
);

-- Tabel untuk deskripsi produk
CREATE TABLE IF NOT EXISTS product_description (
    product_description_id INT AUTO_INCREMENT PRIMARY KEY,
    deskripsi_produk TEXT
);

-- Tabel untuk metode pembayaran
CREATE TABLE IF NOT EXISTS payment_method (
    payment_method_id INT AUTO_INCREMENT PRIMARY KEY,
    nama_metode VARCHAR(50) NOT NULL
);

-- Tabel untuk produk
CREATE TABLE product (
    product_id INT AUTO_INCREMENT PRIMARY KEY,
    nama_produk VARCHAR(255) NOT NULL,
    harga DECIMAL(10, 2) NOT NULL,
    stok INT NOT NULL,
    product_type_id INT,
    operator_id INT,
    product_description_id INT,
    FOREIGN KEY (product_type_id) REFERENCES product_type(product_type_id),
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id),
    FOREIGN KEY (product_description_id) REFERENCES product_description(product_description_id)
);

		-- jawaban no 2 c  
        
-- Tabel untuk transaksi (transaction)
CREATE TABLE transaction (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    tanggal_transaksi TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_harga DECIMAL(10, 2) NOT NULL,
    metode_pembayaran_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (metode_pembayaran_id) REFERENCES payment_method(payment_method_id)
);

-- Tabel untuk detail transaksi (transaction detail)
CREATE TABLE transaction_detail (
    transaction_detail_id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    jumlah_barang INT NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transaction(transaction_id),
    FOREIGN KEY (product_id) REFERENCES product(product_id)
);
 -- jawaban no 3 
 CREATE TABLE kurir (
	id INT AUTO_INCREMENT PRIMARY KEY,
    name varchar(255) NOT NULL, 
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp
    ); 
    
-- jawaban no 4 
	ALTER TABLE kurir
    ADD  ongkos_dasar DECIMAL(10, 2) NOT NULL; 
-- jawaban no 5 
RENAME  TABLE kurir TO shipping;
-- jawaban no 6 
DROP TABLE shipping;
-- jawaban no 7 one to one
 CREATE TABLE payment_method_description ( 
	payment_method_id INT PRIMARY KEY,
    description TEXT, 
    foreign key (payment_method_id) references payment_method (payment_method_id)
    );
-- jawaban no 7 one to many 
CREATE TABLE alamat ( 
	alamat_id  INT AUTO_INCREMENT PRIMARY KEY, 
	user_id INT, 
    alamat varchar (255), 
    foreign key (user_id) REFERENCES user (user_id)
    );
-- jawaban no 7 many to many 
CREATE TABLE  user_payment_method_detail ( 
	user_id INT,
    payment_method_id INT, 
    primary key (user_id, payment_method_id), 
    foreign key (user_id) references user (user_id),
    foreign key (payment_method_id) references payment_method(payment_method_id)
    );
