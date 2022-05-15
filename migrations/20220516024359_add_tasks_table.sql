-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS `tasks` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `created_by` int NOT NULL,
    `title` varchar(255) NOT NULL,
    `status` ENUM('open','inprogress', 'done', 'holding', 'canceled') DEFAULT 'open',
    `deadline` timestamp,
    `assignee` int,
    `description` text,
    `parent_id` int DEFAULT NULL,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS `tasks`;