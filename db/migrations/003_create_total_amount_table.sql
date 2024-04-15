CREATE TABLE IF NOT EXISTS TotalAmount (
    id BIGSERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    remaining_amount INT NOT NULL,
    included_category VARCHAR(255) NOT NULL,
    label VARCHAR(255) NOT NULL,
    statuss VARCHAR(255) NOT NULL
);