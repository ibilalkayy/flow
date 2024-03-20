CREATE TABLE IF NOT EXISTS Alert (
    id BIGSERIAL PRIMARY KEY,
    alert_methods VARCHAR(255) NOT NULL,
    alert_frequencies VARCHAR(255) NOT NULL
);