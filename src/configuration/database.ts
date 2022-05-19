import * as dotenv from "dotenv";
import { TypeOrmModuleOptions } from "@nestjs/typeorm";

function getDatabase(): TypeOrmModuleOptions | any {
  if (!Boolean(process.env.NODE_ENV === "production")) {
    dotenv.config({ path: ".env" });
  }

  return {
    type: process.env.DATABASE_TYPE,
    host: process.env.DATABASE_HOST,
    database: process.env.DATABASE_NAME,
    username: process.env.DATABASE_USER,
    port: process.env.DATABASE_PORT,
    password: process.env.DATABASE_PASSWORD,
    entities: [__dirname + "/../**/*.entity{.ts,.js}"],
    migrations: ["dist/database/migrations/**/*.js"],
    cli: {
      	migrationsDir: "src/database/migrations",
    },
    seeds: ["dist/database/seeds/**/*.seed.js"],
    factories:  ["dist/database/factories/**/*.js"]
  };
}

const databaseConfiguration = getDatabase();

export default databaseConfiguration;
