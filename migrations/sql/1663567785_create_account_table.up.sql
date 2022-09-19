CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER(20) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NULL DEFAULT NULL,
    type VARCHAR(255) NULL DEFAULT NULL,
    description TEXT NULL DEFAULT NULL,
    user_id INTEGER(20) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (id)
);