insert into tasks(id,name,content,user_id)
values
(UUID(), 'test-1', 'content test-1', '1231231'),
(UUID(), 'test-2', 'content test-2', '1231231'),
(UUID(), 'test-3', 'content test-3', '1231232'),
(UUID(), 'test-4', 'content test-4', '1231232'),
(UUID(), 'test-5', 'content test-5', '1231233'),
(UUID(), 'test-6', 'content test-6', '1231233'),
(UUID(), 'test-7', 'content test-7', '1231234'),
(UUID(), 'test-8', 'content test-8', '1231234'),
(UUID(), 'test-9', 'content test-9', '1231235'),
(UUID(), 'test-10', 'content test-10', '1231235');

insert into user_configs(id, user_id, config_type, value)
values
(UUID(), '1231231', 'RATE_LIMIT_ADD_TASK_PER_DAY', '10'),
(UUID(), '1231232', 'RATE_LIMIT_ADD_TASK_PER_DAY', '11'),
(UUID(), '1231233', 'RATE_LIMIT_ADD_TASK_PER_DAY', '5'),
(UUID(), '1231234', 'RATE_LIMIT_ADD_TASK_PER_DAY', '15'),
(UUID(), '1231235', 'RATE_LIMIT_ADD_TASK_PER_DAY', '20'),
(UUID(), '23223332', 'RATE_LIMIT_ADD_TASK_PER_DAY', '10'),
(UUID(), '23223333', 'RATE_LIMIT_ADD_TASK_PER_DAY', '5'),
(UUID(), '23223334', 'RATE_LIMIT_ADD_TASK_PER_DAY', '5'),
(UUID(), '23223335', 'RATE_LIMIT_ADD_TASK_PER_DAY', '7'),
(UUID(), '23223336', 'RATE_LIMIT_ADD_TASK_PER_DAY', '8'),
(UUID(), '23223337', 'RATE_LIMIT_ADD_TASK_PER_DAY', '9');