CREATE TABLE IF NOT EXISTS task (
  pk SERIAL PRIMARY KEY,
  id VARCHAR UNIQUE NOT NULL,
  user_pk INT NOT NULL REFERENCES "user"(pk),
  message VARCHAR NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)