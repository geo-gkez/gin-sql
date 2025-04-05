-- Create customers table
CREATE TABLE customers
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(100)        NOT NULL,
    last_name  VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    phone      VARCHAR(20)
);

-- Create accounts table
CREATE TABLE accounts
(
    id                  SERIAL PRIMARY KEY,
    customer_id         INTEGER            NOT NULL,
    account_number      VARCHAR(50) UNIQUE NOT NULL,
    balance             DECIMAL(15, 2)     NOT NULL DEFAULT 0.00,
    account_description TEXT,
    created_at          TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP,
    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE
);

-- Create index on foreign key for performance
CREATE INDEX idx_accounts_customer_id ON accounts (customer_id);