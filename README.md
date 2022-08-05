# To-go

### Technology
- Hexagonal Architecture
  - Hexagonal Architecture promotes the separation of concerns by encapsulating logic in different layers of the application

- Spring WebFlux
  - Spring Webflux Framework is part of Spring 5 and it provides reactive programming support for web applications
  
### Prerequisites
- Docker
- JDK 11

### Start on local
- **Step 1**: Run docker-compose 
```zsh
docker-compose -p togo-local -f docker-compose.yml up
```

- **Step 2**: Run project with env in file ".env.sample"

### Run unitest locally

- Run 
```zsh
gradle server:test
```

### I can improve
- Use **UUID** to save id replace **String**
- User **redis** to **caching**

### DB Schema

```sql
-- users definition
CREATE TABLE users
(
  id         varchar(64) NOT NULL
    primary key,
  limit_task int         NOT NULL
);

-- tasks definition
CREATE TABLE tasks
(
  id          varchar(64)                        NOT NULL
    primary key,
  title       varchar(64)                        NOT NULL,
  description varchar(255)                       NOT NULL,
  user_id     varchar(64)                        NOT NULL,
  created_at  DateTime DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at  DateTime DEFAULT CURRENT_TIMESTAMP NOT NULL,
  constraint tasks_ibfk_1
    foreign key (user_id) references users (id)
      on delete cascade
);

```