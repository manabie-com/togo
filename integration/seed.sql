INSERT INTO "users" ("id", "password", "max_todo") VALUES
('firstUser', '$2a$10$Fs1rOJV67aVdM/DJHp2F0eJkFVmTKSBBOF0B/70wM0SLGAKk0WMTy', '5');

INSERT INTO "tasks" ("id", "content", "user_id", "created_date") VALUES
('3f44bbd3-8550-4a7c-a654-4495c060c36d', 'content_1', 'firstUser', '2020-08-20'),
('498f276a-9145-4892-a480-f106b4708240', 'content_2', 'firstUser', '2020-08-20'),
('a5a2ad9f-9472-4b02-a57e-659d0d561a0f', 'content_3', 'firstUser', '2020-08-20'),
('af2760a4-da6f-402b-9339-856c287b66a1', 'content_4', 'firstUser', '2020-08-20');