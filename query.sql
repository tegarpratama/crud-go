CREATE DATABASE visiku_test;

use visiku_test;

CREATE TABLE product_categories (
    id CHAR(36) NOT null primary KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO product_categories (id, name) 
VALUES 
(UUID(), 'Computer'),
(UUID(), 'Gadget'),
(UUID(), 'Keyboard');


CREATE TABLE products (
    id CHAR(36) NOT null primary KEY,
    name VARCHAR(255) NOT null,
    description text not null,
    category_id CHAR(36) not null,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);