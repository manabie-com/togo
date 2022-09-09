create table tasks (
    id varchar(40) primary key,
    name varchar(255) not null,
    content varchar(1000),
    user_id varchar(40)
);

create table user_configs(
    id varchar(40) primary key,
    user_id varchar(40),
    config_type varchar(255),
    value varchar(255)
)