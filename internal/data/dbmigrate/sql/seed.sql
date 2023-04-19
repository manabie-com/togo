INSERT INTO users (id, name, email, password_hash, daily_max_todo, date_created, date_updated) VALUES
	('5cf37266-3473-4006-984f-9325122678b7', 'User Gopher', 'user@example.com', '$2a$10$1ggfMVZV6Js0ybvJufLRUOWHS5f6KneuP0XwwHpJ8L8ipdry9f2/a', 1, '2019-03-24 00:00:00', '2019-03-24 00:00:00'),
	('45b5fbd3-755f-4379-8f07-a58d4a30fa2f', 'User 2 Gopher', 'user2@example.com', '$2a$10$9/XASPKBbJKVfCAZKDH.UuhsuALDr5vVm6VrYA9VFR8rccK86C1hW', 2, '2019-03-24 00:00:00', '2019-03-24 00:00:00')
	ON CONFLICT DO NOTHING;
