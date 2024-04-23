CREATE TABLE IF NOT EXISTS Budget (
    id BIGSERIAL PRIMARY KEY,
    categories VARCHAR(255) NOT NULL,
    amounts INT NOT NULL,
    spent INT NOT NULL,
    remaining INT NOT NULL,
    UNIQUE(categories)
);

CREATE TABLE IF NOT EXISTS History (
    id BIGSERIAL PRIMARY KEY,
    dates VARCHAR(255) NOT NULL,
    categories VARCHAR(255) NOT NULL,
    amounts INT NOT NULL,
    transaction_ids VARCHAR(255) NOT NULL,
    blockchains VARCHAR(255) NOT NULL,
    addresses VARCHAR(255) NOT NULL
);