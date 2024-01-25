CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    cnpj VARCHAR(14) NOT NULL
);

CREATE TABLE IF NOT EXISTS product_segments (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    company_id UUID NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    company_id UUID NOT NULL,
    admin BOOLEAN NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    company_id UUID NOT NULL,
    segment_id UUID NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id),
    FOREIGN KEY (segment_id) REFERENCES product_segments(id)
);

CREATE TABLE IF NOT EXISTS suppliers (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    company_id UUID NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY NOT NULL,
    product_id UUID NOT NULL,
    transaction_type VARCHAR(10) NOT NULL,
    quantity INT NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS deposits (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    company_id UUID NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE IF NOT EXISTS deposit_products (
    id UUID PRIMARY KEY NOT NULL,
    product_id UUID NOT NULL,
    deposit_id UUID NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (deposit_id) REFERENCES deposits(id)
);
