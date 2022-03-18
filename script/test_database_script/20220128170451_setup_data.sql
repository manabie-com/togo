-- +goose Up
-- +goose StatementBegin
INSERT INTO users VALUES (1, 'example', '$2a$14$QsCCoXE0O84KFzMMOfMQ8O.szZqS9pRzaxAovr4/WAN0tgO//A9S2', 10);
INSERT INTO users VALUES (2, 'example_1', '$2a$14$sDFKaW4u/xVUf6uqcufz4.haedINBjBTArAQvmWbXO9X4lbrLiEgG', 0);
INSERT INTO tasks VALUES (1, 'Quét nhà 1', 1);
INSERT INTO tasks VALUES (2, 'Quét nhà 2', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
