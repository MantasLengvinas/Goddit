CREATE TABLE IF NOT EXISTS posts(
    id integer PRIMARY KEY AUTOINCREMENT,
    user_id integer NOT NULL,
    content string NOT NULL
);