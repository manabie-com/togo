--liquibase formatted sql

--changeset chiennv:2

INSERT INTO `users` (`id`, `limit_task`)
VALUES ('0xaa91b0c2bc854c619253accc3cfc9270', 2);
INSERT INTO `users` (`id`, `limit_task`)
VALUES ('0xaa91b0c2bf854c619253accc3cfc9270', 1);
INSERT INTO `users` (`id`, `limit_task`)
VALUES ('0xaa91b0c2bc854c698253acdc3cfc9270', 3);

--changeset chiennv:4

INSERT INTO `tasks` (`id`, `title`, `description`, `user_id`, `created_at`, `updated_at`)
VALUES ('0xb7d29119bee54ae9b9defb9c5b7a087d', 'task1', 'ssss', '0xaa91b0c2bc854c619253accc3cfc9270',
        '2022-08-03 13:52:03', '2022-08-03 13:52:03');
INSERT INTO `tasks` (`id`, `title`, `description`, `user_id`, `created_at`, `updated_at`)
VALUES ('0x9864330b3f9f448cbdd8fb59e98c776e', 'task2', 'asfdf', '0xaa91b0c2bf854c619253accc3cfc9270',
        '2022-08-03 13:52:03', '2022-08-03 13:52:03');
INSERT INTO `tasks` (`id`, `title`, `description`, `user_id`, `created_at`, `updated_at`)
VALUES ('0xb7d29119bee54ae9b9defb9c5b7a097d', 'task3', 'aaaaa', '0xaa91b0c2bc854c619253accc3cfc9270',
        '2022-08-03 13:52:03', '2022-08-03 13:52:03');
INSERT INTO `tasks` (`id`, `title`, `description`, `user_id`, `created_at`, `updated_at`)
VALUES ('0xd4b048316694424bbcfb62f51a0b6031', 'task4', 'aaaaa', '0xaa91b0c2bc854c698253acdc3cfc9270',
        '2022-08-03 13:52:03', '2022-08-03 13:52:03');



