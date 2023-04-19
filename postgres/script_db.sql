
CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY,
    user_name VARCHAR(50) UNIQUE NOT NULL,
    pass_word VARCHAR NOT NULL,
    limit_task INT DEFAULT 5 NOT null,
    PRIMARY KEY(id)
);

CREATE TABLE tasks (
    task_id VARCHAR NOT NULL,
    user_id INT NOT NULL REFERENCES users(id),
    "content" VARCHAR(100) NOT NULL,
    created_date DATE NOT NULL,
    event_time DATE NOT NULL,
    PRIMARY KEY (task_id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);