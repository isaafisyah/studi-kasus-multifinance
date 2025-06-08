CREATE TABLE record_transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    konsumen_id INT NOT NULL,
    nomor_kontrak VARCHAR(50) NOT NULL,
    otr VARCHAR(50) NOT NULL,
    admin_fee BIGINT NOT NULL,
    jumlah_cicilan BIGINT NOT NULL,
    jumlah_bunga BIGINT NOT NULL,
    nama_aset VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
