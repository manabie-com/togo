BEGIN;

ALTER TABLE "public"."users" DROP CONSTRAINT user_PK;

DROP TABLE IF EXISTS "public"."users";

COMMIT;
