CREATE TABLE sessions (
    username TEXT PRIMARY KEY,
    session_token VARCHAR(255) NOT NULL UNIQUE,
    expiry DATETIME NOT NULL
);