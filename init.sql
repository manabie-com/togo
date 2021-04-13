/*
 Navicat Premium Data Transfer

 Source Server         : postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 130002
 Source Host           : localhost:5432
 Source Catalog        : manabie
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130002
 File Encoding         : 65001

 Date: 13/04/2021 14:28:15
*/

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS "public"."tasks";
CREATE TABLE "public"."tasks" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "content" text COLLATE "pg_catalog"."default" NOT NULL,
  "user_id" text COLLATE "pg_catalog"."default" NOT NULL,
  "created_date" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."tasks" OWNER TO "my_user";

-- ----------------------------
-- Records of tasks
-- ----------------------------
BEGIN;
INSERT INTO "public"."tasks" VALUES ('e1da0b9b-7ecc-44f9-82ff-4623cc50446a', 'first content', 'firstUser', '2020-06-29');
INSERT INTO "public"."tasks" VALUES ('055261ab-8ba8-49e1-a9e8-e9f725ba9104', 'second content', 'firstUser', '2020-06-29');
INSERT INTO "public"."tasks" VALUES ('2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a', 'another content', 'firstUser', '2020-06-29');
INSERT INTO "public"."tasks" VALUES ('e35e13f8-35f3-409f-8e2f-f3e0173fcca3', 'sadsa', 'firstUser', '2020-08-10');
INSERT INTO "public"."tasks" VALUES ('2a73a4d5-dd05-4c77-bcbd-f5e51a6d6809', 'sadsad', 'firstUser', '2020-08-11');
INSERT INTO "public"."tasks" VALUES ('201f7fee-b512-4785-adef-5f639c887348', 'another content123', 'firstUser', '2021-04-13');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" text COLLATE "pg_catalog"."default" NOT NULL,
  "password" text COLLATE "pg_catalog"."default" NOT NULL,
  "max_todo" int4 NOT NULL
)
;
ALTER TABLE "public"."users" OWNER TO "my_user";

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO "public"."users" VALUES ('firstUser', 'example', 5);
COMMIT;

-- ----------------------------
-- Primary Key structure for table tasks
-- ----------------------------
ALTER TABLE "public"."tasks" ADD CONSTRAINT "tasks_PK" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_PK" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table tasks
-- ----------------------------
ALTER TABLE "public"."tasks" ADD CONSTRAINT "tasks_FK" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
