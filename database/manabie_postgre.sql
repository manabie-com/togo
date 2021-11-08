-- Adminer 4.8.1 PostgreSQL 14.0 (Debian 14.0-1.pgdg110+1) dump

DROP TABLE IF EXISTS "tasks";
DROP SEQUENCE IF EXISTS tasks_task_id_seq;
CREATE SEQUENCE tasks_task_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."tasks" (
    "task_id" integer DEFAULT nextval('tasks_task_id_seq') NOT NULL,
    "user_id" integer DEFAULT '0' NOT NULL,
    "content" character varying(255) NOT NULL,
    "created_date" timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "tasks_pkey" PRIMARY KEY ("task_id")
) WITH (oids = false);


DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_user_id_seq;
CREATE SEQUENCE users_user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."users" (
    "user_id" integer DEFAULT nextval('users_user_id_seq') NOT NULL,
    "username" character varying(50) DEFAULT '' NOT NULL,
    "password" character varying(32) DEFAULT '' NOT NULL,
    "max_todo" integer DEFAULT '5' NOT NULL,
    CONSTRAINT "username" UNIQUE ("username"),
    CONSTRAINT "users_pkey" PRIMARY KEY ("user_id")
) WITH (oids = false);

INSERT INTO "users" ("user_id", "username", "password", "max_todo") VALUES
(1,	'manabie',	'19c4bed05adc55c002d93e6fefe56df2',	5),
(2,	'test',	'b18654785eea30528da6be8a50185428',	6);

ALTER TABLE ONLY "public"."tasks" ADD CONSTRAINT "tasks_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(user_id) NOT DEFERRABLE;

-- 2021-11-03 06:27:07.297467+00
