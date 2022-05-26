CREATE DATABASE todo;

USE todo;

CREATE TABLE IF NOT EXISTS tbl_user
(
    id            int PRIMARY KEY AUTO_INCREMENT,
    username      VARCHAR(50)  NOT NULL,
    password      VARCHAR(255) NOT NULL,
    limit_per_day INT          NOT NULL,
    created_time  DATETIME DEFAULT NOW()
) ENGINE=INNODB;

CREATE INDEX username_ind USING BTREE ON tbl_user(username);


CREATE TABLE IF NOT EXISTS tbl_task
(
    id int PRIMARY KEY AUTO_INCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    created_by int,
    date DATE NOT NULL DEFAULT (curdate()),
    created_time DATETIME DEFAULT NOW(),
    FOREIGN KEY (created_by) REFERENCES tbl_user(id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=INNODB;