--- create user table
CREATE TABLE public.users (
  id bigserial NOT NULL,
  username varchar(50) unique NOT NULL,
  password varchar(50) NOT NULL,
  created_date timestamp(6) NULL DEFAULT LOCALTIMESTAMP,
  max_todo INTEGER DEFAULT 5 NOT NULL,
  CONSTRAINT users_PK PRIMARY KEY (id)
);

--- create tasks table
CREATE TABLE public.tasks (
	id bigserial NOT NULL,
	content TEXT NOT NULL,
	user_id int8 NOT NULL,
  created_date timestamp(6) NULL DEFAULT LOCALTIMESTAMP,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

--- create index
CREATE INDEX "user_id_IDX"  ON public."tasks" USING btree ("user_id");
CREATE INDEX "created_date_IDX"  ON public."tasks" USING btree ("created_date");



--- Insert users
INSERT INTO public.users (id, username, "password", created_date, max_todo) VALUES(1, 'user1', '$2a$10$MSlzbaal5/i3PMaGMDocjefbyQzdR58MWMyWA1JrFScgsmO4Fku62', '2021-10-31 08:20:35.159', 5);
INSERT INTO public.users (id, username, "password", created_date, max_todo) VALUES(2, 'user2', '$2a$10$MSlzbaal5/i3PMaGMDocjefbyQzdR58MWMyWA1JrFScgsmO4Fku62', '2021-10-31 08:20:35.159', 5);


--- Insert tasks
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a developer', 1);
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a design', 1);
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a cooker', 1);
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a police', 1);

INSERT INTO public.tasks ("content", user_id) VALUES('I''m a M3P', 2);
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a Sala', 2);
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a M82', 2);
INSERT INTO public.tasks ("content", user_id) VALUES('I''m a PR7', 2);
