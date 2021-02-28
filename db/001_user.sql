-- users definition

CREATE TABLE users (
   id TEXT NOT NULL PRIMARY KEY,
   password TEXT NOT NULL,
   max_todo INT8 DEFAULT 5 NOT NULL
);
