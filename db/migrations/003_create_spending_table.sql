CREATE TABLE IF NOT EXISTS Spending (
    id BIGSERIAL PRIMARY KEY,
    categories VARCHAR(255) NOT NULL,
    category_amounts VARCHAR(255) NOT NULL,
    spending_amounts VARCHAR(255) NOT NULL,
    amount_exceeded VARCHAR(255) NOT NULL,
    UNIQUE(categories)
);