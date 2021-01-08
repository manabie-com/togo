-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
    id            TEXT PRIMARY KEY,
    password_hash TEXT    NOT NULL,
    max_todo      INTEGER NOT NULL DEFAULT 5
);

INSERT INTO users (id, password_hash, max_todo) VALUES('firstUser', '$2a$10$Px50y37hZA.W4h8t2hvDMeIyenU3kDNWx0NCZpBtUyHHJUbW3e1Uu', 5);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users CASCADE;