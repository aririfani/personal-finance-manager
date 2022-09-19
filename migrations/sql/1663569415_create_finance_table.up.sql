CREATE TABLE IF NOT EXISTS finances (
    id INTEGER(20) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NULL DEFAULT NULL,
    account_id INTEGER(20) NOT NULL,
    amount DECIMAL(12, 2) NULL DEFAULT NULL,
    description TEXT NULL DEFAULT NULL,
    user_id INTEGER(20) NOT NULL,
    type VARCHAR(255) NULL DEFAULT NULL,
    transaction_date DATETIME NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (id)
);