create table users
(
    id       uuid    not null
        constraint users_pk
            primary key,
    username text    not null,
    password text    not null,
    max_todo integer not null
);

create unique index users_username_uindex
    on users (username);

create table tasks
(
    id           uuid        not null
        constraint tasks_pk
            primary key,
    user_id      uuid        not null
        constraint tasks_users_id_fk
            references users,
    content      text        not null,
    created_date timestamptz not null
);
