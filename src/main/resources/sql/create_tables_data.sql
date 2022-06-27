CREATE TABLE user (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    limit_tasks_per_day int(11),
    PRIMARY KEY (id)
);


CREATE TABLE task (
    id int NOT NULL AUTO_INCREMENT,
    description varchar(250) NOT NULL,
    title varchar(50) NOT NULL,
    user_id int,
    created_date Datetime,
    updated_date Datetime,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

insert into `user` (name, limit_tasks_per_day)
values ("userA", 2), ("userB", 5)