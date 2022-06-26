BEGIN;

INSERT INTO "public"."users"  (id, username, password, max_task_per_day) VALUES('firstUser','manabie' ,'example', 5);
INSERT INTO "public"."users"  (id, username, password, max_task_per_day) VALUES('secondUser','manabie-test-1' ,'example', 0);

COMMIT;
