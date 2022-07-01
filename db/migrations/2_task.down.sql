BEGIN;

ALTER TABLE "public"."tasks" DROP CONSTRAINT task_FK;
ALTER TABLE "public"."tasks" DROP CONSTRAINT task_PK;

DROP TABLE IF EXISTS "public"."tasks";

COMMIT;
