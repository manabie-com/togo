import { ConnectionOptions } from "typeorm";
import { SnakeNamingStrategy } from "typeorm-naming-strategies";

if (process.env.NODE_ENV === "test") {
  process.env.DB_NAME = "todo_db_test";
}

const config: ConnectionOptions = {
  type: "mysql",
  host: process.env.DB_HOST || "localhost",
  username: process.env.DB_USER || "root",
  password: process.env.DB_PASSWORD || "",
  database: process.env.DB_NAME || "todo_db",
  synchronize: false,
  logging: true,
  entities: ["entity/**/*.ts"],
  migrations: ["migration/**/*.ts"],
  subscribers: ["subscriber/**/*.ts"],
  cli: {
    entitiesDir: "entity",
    migrationsDir: "migration",
    subscribersDir: "subscriber",
  },
  namingStrategy: new SnakeNamingStrategy(),
};

export default config;
