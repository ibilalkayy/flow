CREATE TABLE IF NOT EXISTS Alert (
    id BIGSERIAL PRIMARY KEY,
    categories VARCHAR(255) NOT NULL,
    alert_methods VARCHAR(255) NOT NULL,
    alert_frequencies VARCHAR(255) NOT NULL,
    alert_days INT NOT NULL,
    alert_weekdays VARCHAR(255) NOT NULL,
    alert_hours INT NOT NULL,
    alert_minutes INT NOT NULL,
    UNIQUE(categories)
);