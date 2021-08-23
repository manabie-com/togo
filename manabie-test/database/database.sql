SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users`
(
    `id`       varchar(500)      NOT NULL,
    UNIQUE KEY unique_id (id),
    `password` text              NOT NULL,
    `max_todo` int(50) DEFAULT 5 NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` (`id`, `password`, `max_todo`)
VALUES ('kiennt',
        'c70b5dd9ebfb6f51d09d4132b7170c9d20750a7852f00680f65658f0310e810056e6763c34c9a00b0e940076f54495c169fc2302cceb312039271c43469507dc',
        3);
INSERT INTO `users` (`id`, `password`, `max_todo`)
VALUES ('admin',
        'c70b5dd9ebfb6f51d09d4132b7170c9d20750a7852f00680f65658f0310e810056e6763c34c9a00b0e940076f54495c169fc2302cceb312039271c43469507dc',
        5);

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS `tasks`;
CREATE TABLE IF NOT EXISTS `tasks`
(
    `id`           varchar(500) NOT NULL,
    UNIQUE KEY unique_id (id),
    `content`      text         NOT NULL,
    `created_date` DATETIME     NOT NULL,
    `user_id`      varchar(255) NOT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (`id`),
    CONSTRAINT `tasks_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;

-- ----------------------------
-- Records of tasks
-- ----------------------------
INSERT INTO `tasks` (`id`, `content`, `created_date`, `user_id`)
VALUES ('firstTask', 'Task to go 1', '2021-08-20 00:00:00', 'kiennt');
INSERT INTO `tasks` (`id`, `content`, `created_date`, `user_id`)
VALUES ('secondTask', 'Task to go 2', '2021-08-20 00:00:00', 'admin');

SET FOREIGN_KEY_CHECKS = 1;
