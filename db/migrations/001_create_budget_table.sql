CREATE TABLE IF NOT EXISTS Budget (
    id BIGSERIAL PRIMARY KEY,
    categories VARCHAR(200) NOT NULL,
    amounts INT NOT NULL,
    spent INT NOT NULL,
    remaining INT NOT NULL,
    UNIQUE(categories)
);