CREATE TABLE IF NOT EXISTS users
(
    `id`         BIGINT AUTO_INCREMENT,
    `email`      VARCHAR(128) NOT NULL,
    `name`       VARCHAR(128) NOT NULL,
    `gender`     TINYINT(4)   NOT NULL DEFAULT 0 COMMENT '0: Female, 1: Male, 2: Other',
    `password`   VARCHAR(255) NOT NULL,
    `max_todo`   INT          NOT NULL DEFAULT 10,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
