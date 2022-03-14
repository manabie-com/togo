CREATE TABLE users (
  id INT UNSIGNED AUTO_INCREMENT,
  full_name NVARCHAR(100) NOT NULL,
  daily_record_limit INT UNSIGNED,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);
INSERT INTO users(id, full_name, daily_record_limit) VALUES(1, "Alice", 5);
INSERT INTO users(id, full_name, daily_record_limit) VALUES(2, "Bob", 10);
INSERT INTO users(id, full_name, daily_record_limit) VALUES(3, "Charlie", 15);

CREATE TABLE user_todo_items (
  id INT UNSIGNED AUTO_INCREMENT,
  content NVARCHAR(1000) NOT NULL,
  is_done BOOL DEFAULT 0,
  user_id INT UNSIGNED,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

