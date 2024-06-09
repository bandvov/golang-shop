CREATE TABLE products (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    price NUMERIC
);

CREATE TABLE users (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100)
);

-- Add tables for carts and orders as needed
