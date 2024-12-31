CREATE TABLE IF NOT EXISTS fruit_vegetables (
    fruit_vegetable_name TEXT PRIMARY KEY NOT NULL,
    category TEXT CHECK(category IN ('fruit', 'vegetable')) NOT NULL,
    region_name TEXT,
    image_url TEXT,
    FOREIGN KEY (region_name) REFERENCES regions(region_name)
);
