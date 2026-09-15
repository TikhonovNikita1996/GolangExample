-- +goose Up
CREATE TABLE orders (
    order_uuid UUID PRIMARY KEY,
    user_uuid UUID NOT NULL,
    parts_uuid UUID[] NOT NULL,
    total_price REAL NOT NULL,
    transaction_uuid UUID NULL,
    payment_method TEXT NULL,
    status TEXT NOT NULL
);
-- +goose Down
DROP TABLE orders;
