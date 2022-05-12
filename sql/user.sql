CREATE TABLE IF NOT EXISTS `users`
(
    `id`         BIGINT PRIMARY KEY AUTO_INCREMENT,
    `username`   VARCHAR(50)  NOT NULL UNIQUE,
    `password`   VARCHAR(72)  NOT NULL,
    `task_quote` INT UNSIGNED NOT NULL,
    `is_delete`  BOOLEAN      NOT NULL DEFAULT FALSE
);

ALTER TABLE `users`
    ADD INDEX `user_username_idx` (`username`);


