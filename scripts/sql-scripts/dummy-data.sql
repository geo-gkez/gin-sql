-- Insert sample customers
INSERT INTO customers (first_name, last_name, email, phone)
VALUES
    ('John', 'Doe', 'john.doe@example.com', '555-1234'),
    ('Jane', 'Smith', 'jane.smith@example.com', '555-5678'),
    ('George', 'Gkezeris', 'gg@gmail.com', '699999');

-- Insert sample accounts
INSERT INTO accounts (customer_id, account_number, balance, account_description)
VALUES
    (1, 'ACC-10001', 5000.00, 'Checking Account'),
    (1, 'ACC-10002', 25000.00, 'Savings Account'),
    (2, 'ACC-20001', 7500.50, 'Checking Account'),
    (2, 'ACC-20002', 42000.75, 'Investment Account'),
    (3, 'ACC-30001', 1500.25, 'Checking Account'),
    (3, 'ACC-30002', 15000.00, 'Savings Account'),
    (3, 'ACC-30003', 50000.00, 'Retirement Account');

-- Example of updating an account
UPDATE accounts
SET balance = balance + 1000.00, updated_at = CURRENT_TIMESTAMP
WHERE account_number = 'ACC-10001';

-- Example of soft-deleting an account
UPDATE accounts
SET deleted_at = CURRENT_TIMESTAMP
WHERE account_number = 'ACC-20002';