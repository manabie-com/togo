import "reflect-metadata";
import { Container } from "inversify";
import { Logger } from "./logger";
import { CacheClient } from "./cache.client";
import { MongoClient } from "./mongo.client";
import { PostgresClient } from "./postgres.client";

export const CORES = {
    Logger: Symbol.for("Logger"),
    CacheClient: Symbol.for("CacheClient"),
    MongoClient: Symbol.for("MongoClient"),
    PostgresClient: Symbol.for("PostgresClient"),
}

export class CoreContainer {
    public static Load(container: Container) {
        container.bind<Logger>(CORES.Logger).to(Logger);
        container.bind<CacheClient>(CORES.CacheClient).to(CacheClient);
        container.bind<MongoClient>(CORES.MongoClient).to(MongoClient);
        container.bind<PostgresClient>(CORES.PostgresClient).to(PostgresClient);
    }
}

export * from "./error.interface";
export * from "./request.interface";
export * from "./api.response";
export * from "./app.container";
export * from "./mongo.client";
export * from "./base.middleware";
export * from "./logger";
export * from "./postgres.client";
