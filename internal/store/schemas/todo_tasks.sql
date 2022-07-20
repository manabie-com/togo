CREATE TABLE IF NOT EXISTS `todo_tasks` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `user_id` bigint(20) UNSIGNED NOT NULL,
    `task_name` VARCHAR(255) NOT NULL
);