create table if not exists users
(
    id       varchar primary key,
    hash     varchar,
    max_todo int
);

create table if not exists tasks
(
    id           uuid primary key,
    content      text,
    user_id      varchar,
    created_date date,
    done         bool,
    deleted      bool
);

truncate users;

truncate tasks;
