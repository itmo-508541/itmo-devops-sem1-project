DROP TABLE prices IF EXISTS;

CREATE TABLE IF NOT EXISTS prices (
    uuid TEXT PRIMARY KEY,
    id INTEGER,
    name TEXT,
    category TEXT,
    price DECIMAL(10, 2),
    create_date DATE,
    group_uuid TEXT
);

CREATE TABLE IF NOT EXISTS reports (
    uuid TEXT PRIMARY KEY,
    id INTEGER,
    name TEXT,
    category TEXT,
    price DECIMAL(10, 2),
    create_date DATE
);
