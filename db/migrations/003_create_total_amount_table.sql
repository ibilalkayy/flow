CREATE TABLE IF NOT EXISTS TotalAmount (
    id BIGSERIAL PRIMARY KEY,
    amount INT NOT NULL,
    included_category VARCHAR(255) NOT NULL,
    excluded_category VARCHAR(255) NOT NULL,
    label VARCHAR(255) NOT NULL,
    statuss VARCHAR(255) NOT NULL
);