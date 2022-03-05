CREATE TABLE IF NOT EXISTS todos
(
    `id`         BIGINT AUTO_INCREMENT,
    `user_id`    BIGINT       NOT NULL,
    `title`      VARCHAR(128) NOT NULL,
    `content`    TEXT         NOT NULL,
    `status`     TINYINT(4)   NOT NULL DEFAULT 0 COMMENT '0: Open, 1: InProgress, 2: Resolved',
    `edited_at`  TIMESTAMP             DEFAULT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP             DEFAULT NULL,

    PRIMARY KEY (`id`),
    CONSTRAINT `todos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
