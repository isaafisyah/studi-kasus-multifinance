CREATE TABLE limits(
    id INT AUTO_INCREMENT PRIMARY KEY,
    konsumen_id INT NOT NULL,
    tenor TINYINT NOT NULL CHECK (tenor IN (1, 2, 3, 4)),
    limit_amount BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT unique_konsumen_tenor UNIQUE (konsumen_id, tenor)
)