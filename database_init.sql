SELECT 'CREATE DATABASE todo'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'todo');\gexec

SELECT 'CREATE DATABASE test'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'test');\gexec

