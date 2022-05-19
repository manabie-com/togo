CREATE TABLE "users" (
                        "id" SERIAL PRIMARY KEY,
                        "user_name" varchar NOT NULL,
                        "hashed_password" varchar NOT NULL,
                        "created_at" timestamp NOT NULL DEFAULT (now()),
                        "updated_at" timestamp NOT NULL DEFAULT (now()),
                        "maximum_task_in_day"  int NOT NULL
);

CREATE TABLE "tasks" (
                        "id" SERIAL PRIMARY KEY,
                        "title" varchar NOT NULL,
                        "user_id" int NOT NULL,
                        "created_at" timestamp NOT NULL DEFAULT (now()),
                        "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");