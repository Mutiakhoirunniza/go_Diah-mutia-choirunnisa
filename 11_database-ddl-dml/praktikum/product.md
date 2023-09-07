use alta_online_shop;
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
