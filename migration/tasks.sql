CREATE TABLE `task` (
  `id` varchar(255) PRIMARY KEY,
  `user_id` varchar(255) NOT NULL,
  `content` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL DEFAULT "pending",
  `created_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `user` (
  `id` varchar(255) PRIMARY KEY,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `task` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
INSERT INTO `user` (`id`, `username`, `password`) 
VALUES ('1', 'user1', '12345'), ('2', 'user2', '12345'), ('3', 'user3', '12345'), ('4', 'user4', '12345'), ('5', 'user5', '12345'), ('6', 'user6', '12345');
