DROP TABLE IF EXISTS prices;
DROP TABLE IF EXISTS reports;

CREATE TABLE IF NOT EXISTS prices (
    serial_id SERIAL PRIMARY KEY,
    id INTEGER,
    name TEXT,
    category TEXT,
    price DECIMAL(10, 2),
    create_date DATE
);
