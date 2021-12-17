CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR NOT NULL,
  detail VARCHAR,
  due_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  reporter_id INTEGER,
  assignee_id INTEGER,
  FOREIGN KEY (reporter_id) REFERENCES users (id),
  FOREIGN KEY (assignee_id) REFERENCES users (id)
)