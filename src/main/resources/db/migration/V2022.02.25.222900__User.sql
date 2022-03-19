CREATE TABLE users (
   id INT PRIMARY KEY AUTO_INCREMENT,
   username varchar(255) NOT NULL,
   password varchar(255) NOT NULL,
   maximum_tasks INT8 DEFAULT 5 NOT NULL
);