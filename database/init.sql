SELECT 'CREATE DATABASE manabie'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'manabie')\gexec