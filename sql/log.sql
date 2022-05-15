CREATE TABLE IF NOT EXISTS `log`
(
    `id`       BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id`  BIGINT       NOT NULL,
    `action`   VARCHAR(256) NOT NULL,
    `datetime` DATETIME     NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE `log`
    ADD INDEX `log_userid_idx` (`user_id`);
