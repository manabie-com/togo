CREATE TABLE IF NOT EXISTS "user" (
  pk SERIAL PRIMARY KEY,
  id VARCHAR UNIQUE NOT NULL,
  task_daily_limit INT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)