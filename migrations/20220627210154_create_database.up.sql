CREATE TABLE IF NOT EXISTS users
(
    id         bigserial not null unique,
    name       varchar(255),
    max_todo   int,
    created_at timestamp default (now() at time zone 'UTC'),
    updated_at timestamp default (now() at time zone 'UTC'),
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS todos
(
    id         bigserial not null,
    name       varchar(255),
    content    varchar(255),
    user_id    bigint,
    created_at timestamp default (now() at time zone 'UTC'),
    updated_at timestamp default (now() at time zone 'UTC'),
    deleted_at timestamp,
    constraint fk_user_id foreign key (user_id) references users (id) on delete cascade
);