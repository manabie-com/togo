CREATE TABLE IF NOT EXISTS `task`
(
    `id`               BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id`          BIGINT       NOT NULL,
    `title`            VARCHAR(256) NOT NULL,
    `description`      TEXT,
    `datetime_created` DATETIME     NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    `datetime_edited`  DATETIME     NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    `is_delete`        BOOLEAN      NOT NULL DEFAULT FALSE
);

ALTER TABLE `task`
    ADD INDEX `task_search_idx` (`user_id`, `datetime_created`);