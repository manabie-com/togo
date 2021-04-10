### Overview
To make it run:
- `docker-compose -f build/docker-compose.yml up --build`
- Import Postman collection from `docs` to check example

### Description
- I apply DDD approach which focus on domain.

- I choose CQRS pattern . It seperate into write side and read side to independent scaling, avoid side-effect, optimized data schemas and simpler queries:
    
    + Write side: change state, update data by command and event without response data.
    + Read side: query data from database without change it.
    
- To save system state, I use Event-sourcing:
    + The stored event not only describes the current state of the object, but also shows the history of its creation.
    + System state can be rebuilt at any time for any state in the past by re-build events.
    + ...
    
- I make code `loose coupling` by Dependency Injection.

- About pagination for task list, I use `cursor` technique which only go to next page or previous page from current page. I generate a cursor point to the first row of the next page when user request a page and respond it to client. It's a trade-off between performance and usefulness. 

### Improve
- Use kafka to publish and consume events.
- Use Websocket for asynchronous.
- Implement rate limiting by `leaky bucket`.
- ...

### DB Schema
```sql
-- users definition

create table if not exists users
(
    id            uuid,
    username      text                     not null,
    password      text                     not null,
    created_at    timestamp with time zone not null,
    updated_at    timestamp with time zone not null,
    old_passwords text[],
    constraint users_pk primary key(id)
);

-- (password is 123456789)
insert into users (id, username, password, created_at, updated_at, old_passwords)
values ('456020ea-257c-4066-8b46-b5b186b2335d','zahj',
        '$2a$04$/wqpIyCXr1.61TAOBdexuuXXNj1ubkNG4XpLOCZF35/Ftfnt4PQri','2021-04-10 14:51:18.095134','2021-04-10 14:51:18.095134',null);

-- user_configs defination

create table if not exists user_configs
(
    id         serial,
    user_id    uuid not null,
    task_limit int,
    is_active  boolean,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    constraint user_configs_pk primary key (id)
);

insert into user_configs (user_id, task_limit, is_active, created_at, updated_at)
values ('456020ea-257c-4066-8b46-b5b186b2335d','30',true,'2021-04-10 14:51:18.095134','2021-04-10 14:51:18.095134');

-- user_tasks defination

create table if not exists user_tasks
(
    id          uuid,
    version     bigint                   not null,
    num_of_tasks int                      not null,
    created_at  timestamp with time zone not null,
    updated_at  timestamp with time zone not null,
    constraint user_tasks_pk primary key (id)
);

-- tasks definition

create table if not exists tasks
(
    id         uuid,
    user_id    uuid,
    content    text                     not null,
    created_at timestamp with time zone not null,
    constraint tasks_pk primary key (id)
);

-- aggregates defination

create table if not exists aggregates
(
	namespace varchar(250) not null,
	aggregate_id uuid not null,
	version bigint,
	constraint aggregates_pkey primary key (namespace, aggregate_id)
);

-- event_store defination

create table if not exists event_store
(
	event_id uuid not null,
	namespace varchar(250),
	aggregate_id uuid,
	aggregate_type varchar(250),
	event_type varchar(250),
	data jsonb,
	timestamp timestamp with time zone,
	version bigint,
	context jsonb,
	constraint event_store_pkey primary key (event_id)
);


```

### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
