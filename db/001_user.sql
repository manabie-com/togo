-- users definition

CREATE TABLE users (
   id TEXT NOT NULL PRIMARY KEY,
   password TEXT NOT NULL,
   max_todo INT8 DEFAULT 5 NOT NULL
);

INSERT INTO users (id, password, max_todo) VALUES('00001', 'example', 5);
