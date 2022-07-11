DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS user_daily_limits;

CREATE TABLE IF NOT EXISTS "tasks" (
        "id" int AUTO_INCREMENT PRIMARY KEY,
        "user_id" int,
        "name" VARCHAR ,
        "description" VARCHAR,
        "target_date" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user_daily_limits" (
        "user_id" int PRIMARY KEY,
        "daily_task_limit" int
);
