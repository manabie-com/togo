-- Dummy user data
INSERT INTO users(id, "name", limit_config)
VALUES (1, 'uuhnaut69', 10)
ON CONFLICT (id) DO NOTHING;