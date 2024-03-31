CREATE TABLE IF NOT EXISTS Alert (
    id BIGSERIAL PRIMARY KEY,
    categories VARCHAR(255) NOT NULL,
    category_amounts VARCHAR(255) NOT NULL,
    alert_methods VARCHAR(255) NOT NULL,
    alert_frequencies VARCHAR(255) NOT NULL,
    alert_days VARCHAR(255) NOT NULL,
    alert_weekdays VARCHAR(255) NOT NULL,
    alert_hours VARCHAR(255) NOT NULL,
    alert_minutes VARCHAR(255) NOT NULL,
    alert_seconds VARCHAR(255) NOT NULL,
    UNIQUE(categories)
);