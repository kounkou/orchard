CREATE TABLE IF NOT EXISTS account_fruit_vegetables (
    account_name TEXT NOT NULL,
    fruit_vegetable_name TEXT NOT NULL,
    PRIMARY KEY (account_name, fruit_vegetable_name),
    FOREIGN KEY (account_name) REFERENCES accounts(username),
    FOREIGN KEY (fruit_vegetable_name) REFERENCES fruit_vegetables(fruit_vegetable_name)
);