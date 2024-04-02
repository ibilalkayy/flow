CREATE TABLE IF NOT EXISTS Budget (
    id BIGSERIAL PRIMARY KEY,
    categories VARCHAR(255) NOT NULL,
    amounts VARCHAR(255) NOT NULL,
    spent VARCHAR(255) NOT NULL,
    remaining VARCHAR(255) NOT NULL,
    UNIQUE(categories)
);