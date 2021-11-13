#### Requirement
- Docker
- node latest
- yarn latest

#### Installation
```
    $ git clone https://github.com/Hoangtt4696/togo.git
    $ yarn setup 
    (setup posgreql by docker)
    $ yarn
    (install packages)
    $ yarn start
    (start server with port 5050)
```

#### Unit Test
```
    $ yarn test
```
```
    $ yarn test $file
    (run specific)
```

#### Document api
```
    Run http://127.0.0.1:5050/api/v1/docs (swagger)
```
OR
```
    Import api in /docs/togo.postman_collection.json (postman)
```
#### DB Schema
```sql
-- users definition

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  max_todo INTEGER DEFAULT 5 NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- username = admin, password = admin
INSERT INTO users (username, password, max_todo) VALUES('admin', '$2b$08$QEY11Qebo9Ss..ed9cYhieSdi3xLy/QFl4NMKMfLcBazqSNmhKteS', 5);

-- tasks definition

CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
