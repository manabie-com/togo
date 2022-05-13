CREATE TABLE IF NOT EXISTS users (
   user_id varchar(255) not null,
   name varchar(255),
   task_per_day int,
   created_at timestamp DEFAULT now(),
   updated_at timestamp DEFAULT now(),
   CONSTRAINT users_PK PRIMARY KEY (user_id)
);

INSERT INTO users (user_id, name, task_per_day) VALUES 
('USER_1', 'Juan Dela Cruz', 3),
('USER_2', 'Pedro De Coco', 5);
('USER_3', 'LeBron James', 10);

CREATE TABLE IF NOT EXISTS tasks (
   id UUID not null,
   assigned_to varchar(255),
   title varchar(255),
   description varchar(255),
   created_at timestamp DEFAULT now(),
   updated_at timestamp DEFAULT now(),
   CONSTRAINT tasks_PK PRIMARY KEY (id),
   CONSTRAINT users_id_FKEY FOREIGN KEY (assigned_to) REFERENCES users(user_id)
);