-- tasks definition

CREATE TABLE tasks (
   id TEXT NOT NULL PRIMARY KEY,
   content TEXT NOT NULL,
   user_id TEXT NOT NULL REFERENCES users(id),
   created_date TEXT NOT NULL
);
