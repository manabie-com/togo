/*
 Navicat Premium Data Transfer

 Source Server         : 2Mau-local
 Source Server Type    : PostgreSQL
 Source Server Version : 130003
 Source Host           : localhost:5433
 Source Catalog        : togo
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 130003
 File Encoding         : 65001

 Date: 09/07/2022 21:01:13
*/


-- ----------------------------
-- Sequence structure for configs_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."configs_id_seq";
CREATE SEQUENCE "public"."configs_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."roles_id_seq";
CREATE SEQUENCE "public"."roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for task_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."task_id_seq";
CREATE SEQUENCE "public"."task_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for configs
-- ----------------------------
DROP TABLE IF EXISTS "public"."configs";
CREATE TABLE "public"."configs" (
  "id" int4 NOT NULL DEFAULT nextval('configs_id_seq'::regclass),
  "role" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "limit" int4 NOT NULL,
  "createdAt" timestamptz(6) NOT NULL,
  "updatedAt" timestamptz(6) NOT NULL
)
;

-- ----------------------------
-- Records of configs
-- ----------------------------
INSERT INTO "public"."configs" VALUES (9, 'Admin', 100, '2022-07-07 03:47:49.977+00', '2022-07-07 03:47:49.977+00');

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."roles";
CREATE TABLE "public"."roles" (
  "id" int4 NOT NULL DEFAULT nextval('roles_id_seq'::regclass),
  "code" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "status" bool NOT NULL,
  "createdAt" timestamptz(6) NOT NULL,
  "updatedAt" timestamptz(6) NOT NULL
)
;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO "public"."roles" VALUES (9, 'Admin', 't', '2022-07-07 03:47:49.981+00', '2022-07-07 03:47:49.981+00');

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS "public"."task";
CREATE TABLE "public"."task" (
  "id" int4 NOT NULL DEFAULT nextval('task_id_seq'::regclass),
  "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "description" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "text" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "author" int4 NOT NULL,
  "createdAt" timestamptz(6) NOT NULL,
  "updatedAt" timestamptz(6) NOT NULL
)
;

-- ----------------------------
-- Records of task
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "role" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "createdAt" timestamptz(6) NOT NULL,
  "updatedAt" timestamptz(6) NOT NULL
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (2, 'Đình Bửu', 'dinhbuu1208@gmail.com', '$2a$08$FooFv5KZCP7.pmTAOWqrH.Cqa1TVS434GdWBILdukVzje64/6Ad9O', 'Admin', '2022-07-06 18:56:37.598+00', '2022-07-06 18:56:37.598+00');

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."configs_id_seq"
OWNED BY "public"."configs"."id";
SELECT setval('"public"."configs_id_seq"', 20, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."roles_id_seq"
OWNED BY "public"."roles"."id";
SELECT setval('"public"."roles_id_seq"', 20, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."task_id_seq"
OWNED BY "public"."task"."id";
SELECT setval('"public"."task_id_seq"', 115, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."users"."id";
SELECT setval('"public"."users_id_seq"', 46, true);

-- ----------------------------
-- Primary Key structure for table configs
-- ----------------------------
ALTER TABLE "public"."configs" ADD CONSTRAINT "configs_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table configs
-- ----------------------------
ALTER TABLE "public"."configs" ADD CONSTRAINT "roles_role_key" UNIQUE ("role");

-- ----------------------------
-- Primary Key structure for table roles
-- ----------------------------
ALTER TABLE "public"."roles" ADD CONSTRAINT "roles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table task
-- ----------------------------
ALTER TABLE "public"."task" ADD CONSTRAINT "task_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_email_key" UNIQUE ("email");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
