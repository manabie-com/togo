DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id         int(11) NOT NULL AUTO_INCREMENT,
    login_id   varchar(50) NOT NULL,
    password   varchar(50) NOT NULL,
    status     int(11) NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY UK_login_id (login_id)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO users (login_id, password) VALUE ('test', 'test');

DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks
(
    id         int(11) NOT NULL AUTO_INCREMENT,
    user_id    int(11) NOT NULL,
    content    text,
    status     int(11) NOT NULL DEFAULT '1',
    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
