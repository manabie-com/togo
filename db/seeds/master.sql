DELETE FROM users WHERE id = 1;
INSERT INTO users (id, username, password) VALUES (1, 'linh', crypt('linhdeptrai', gen_salt('bf', 8)));