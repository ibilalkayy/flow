CREATE TABLE IF NOT EXISTS Alert (
    id BIGSERIAL PRIMARY KEY,
    totals VARCHAR(255) NOT NULL,
    categories VARCHAR(255) NOT NULL,
    alert_methods VARCHAR(255) NOT NULL,
    alert_frequencies VARCHAR(255) NOT NULL,
    UNIQUE(categories)
);