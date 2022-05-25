SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS `users` (
   `id` int PRIMARY KEY AUTO_INCREMENT,
   `email` varchar(50) UNIQUE NOT NULL,
    `password` varchar(50) NOT NULL,
    `salt` varchar(50) NOT NULL,
    `daily_task_limit` int NOT NULL DEFAULT 5,
    `status` smallint unsigned NOT NULL DEFAULT 1,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB;

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
