CREATE DATABASE togo_nodejs;

\c togo_nodejs;

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  max_todo INTEGER DEFAULT 5 NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO users (username, password, max_todo) VALUES('admin', '$2b$08$QEY11Qebo9Ss..ed9cYhieSdi3xLy/QFl4NMKMfLcBazqSNmhKteS', 5);

CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);