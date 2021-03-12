
BEGIN TRANSACTION;
CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);
-- firstUser hashed with 10 rounds
INSERT INTO users VALUES('firstUser','$2y$10$nxL74V8dU1F7DAz9xfaQvuZTJ27ZmUNa4sucT2bV1Vtq.MI8EbX/a',5);
CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
INSERT INTO tasks VALUES('e1da0b9b-7ecc-44f9-82ff-4623cc50446a','first content','firstUser','2020-06-29');
INSERT INTO tasks VALUES('055261ab-8ba8-49e1-a9e8-e9f725ba9104','second content','firstUser','2020-06-29');
INSERT INTO tasks VALUES('2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a','another content','firstUser','2020-06-29');
INSERT INTO tasks VALUES('e35e13f8-35f3-409f-8e2f-f3e0173fcca3','fourth content','firstUser','2020-08-10');
INSERT INTO tasks VALUES('2a73a4d5-dd05-4c77-bcbd-f5e51a6d6809','fifth content','firstUser','2020-08-11');
COMMIT;
