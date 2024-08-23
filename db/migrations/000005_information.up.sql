CREATE TABLE IF NOT EXISTS information (
    information_name TEXT PRIMARY KEY NOT NULL,
    description TEXT,
    fruit_vegetable_name TEXT NOT NULL,
    FOREIGN KEY (fruit_vegetable_name) REFERENCES fruit_vegetables(fruit_vegetable_name)
);