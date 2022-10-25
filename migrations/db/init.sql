-- +migrate Up

create table tasks (
    id serial primary key,
    content text,
    user_id varchar(36),
    status smallint default 1,
    created_at timestamp without time zone default current_timestamp,
    updated_at timestamp without time zone default current_timestamp,
    deleted_at timestamp without time zone default NULL
);

create index idx_tasks_user_id_status on tasks (user_id,status);