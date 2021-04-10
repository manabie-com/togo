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

create table if not exists user_tasks
(
    id          uuid,
    version     bigint                   not null,
    num_of_tasks int                      not null,
    created_at  timestamp with time zone not null,
    updated_at  timestamp with time zone not null,
    constraint user_tasks_pk primary key (id)
);

create table if not exists tasks
(
    id         uuid,
    user_id    uuid,
    content    text                     not null,
    created_at timestamp with time zone not null,
    constraint tasks_pk primary key (id)
);

create table if not exists aggregates
(
	namespace varchar(250) not null,
	aggregate_id uuid not null,
	version bigint,
	constraint aggregates_pkey primary key (namespace, aggregate_id)
);

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

insert into users (id, username, password, created_at, updated_at, old_passwords)
values ('456020ea-257c-4066-8b46-b5b186b2335d','zahj',
        '$2a$04$/wqpIyCXr1.61TAOBdexuuXXNj1ubkNG4XpLOCZF35/Ftfnt4PQri','2021-04-10 14:51:18.095134','2021-04-10 14:51:18.095134',null);

insert into user_configs (user_id, task_limit, is_active, created_at, updated_at)
values ('456020ea-257c-4066-8b46-b5b186b2335d','30',true,'2021-04-10 14:51:18.095134','2021-04-10 14:51:18.095134');
