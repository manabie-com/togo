-- +goose Up
-- +goose StatementBegin
-- Time zone +7
SET TIME ZONE 'Asia/Ho_Chi_Minh';

CREATE TABLE users(
    id smallserial PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL,
    limit_task_per_day smallint NOT NULL,
    is_deleted boolean DEFAULT FALSE
);

INSERT INTO users
    (id, name, limit_task_per_day, is_deleted)
VALUES
    (1, 'John', 5, FALSE),
    (2, 'Jack', 1, FALSE),
    (3, 'Mike', 3, TRUE);

CREATE TABLE tasks(
    id smallserial PRIMARY KEY,
    content text,
    user_id smallint,
    date_assign timestamp,
    is_deleted boolean DEFAULT FALSE,
    CONSTRAINT fk_tasks_users
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
);

INSERT INTO tasks
    (id, content, user_id, is_deleted)
VALUES
    (1, 'Task 1', 1, FALSE),
    (2, 'Task 2', 1, FALSE),
    (3, 'Task 3', 2, FALSE),
    (4, 'Task 4', 3, FALSE),
    (5, 'Task 5', 1, TRUE),
    (6, 'Task 6', NULL, FALSE),
    (7, 'Task 7', NULL, FALSE),
    (8, 'Task 8', NULL, FALSE),
    (9, 'Task 9', NULL, FALSE);
UPDATE tasks
    SET date_assign = NOW()
    WHERE user_id IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks, users;
-- +goose StatementEnd
