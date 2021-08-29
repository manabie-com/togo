import { injectable } from "inversify";
import dotenv from "dotenv";
import * as logger from "loglevel";
import { setLogLevel } from "@typegoose/typegoose";
import { Pool, PoolClient } from 'pg';

@injectable()
export class PostgresClient {

    private _pool: Pool;
    private readonly _url!: string;
    private readonly _hideUrl!: string;

    constructor() {
        dotenv.config();
        const DB_USER = process.env.DB_USER;
        const DB_PASS = process.env.DB_PASS;
        const DB_HOST = process.env.DB_HOST;
        const DB_PORT = process.env.DB_PORT || 80;
        const DB_NAME = process.env.DB_NAME;
        const DB_LOG_LEVEL = process.env.DB_LOG_LEVEL as logger.LogLevelDesc || 'warn';

        setLogLevel(DB_LOG_LEVEL);
        this._url = `postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}`;
        this._hideUrl = `postgres://${DB_USER}:xxx@${DB_HOST}:${DB_PORT}/${DB_NAME}`;
        if (!DB_USER || !DB_PASS) {
            this._url = `postgres://${DB_HOST}:${DB_PORT}/${DB_NAME}`;
            this._hideUrl = `postgres://${DB_HOST}:${DB_PORT}/${DB_NAME}`;
        }

        this._pool = new Pool({max: 20, connectionString: this._url, idleTimeoutMillis: 30000});
    }

    public static newInstance(): PostgresClient {
        return new PostgresClient();
    }

    public get pool(): Pool {
        if (!this._pool) {
            this._pool = new Pool({max: 20, connectionString: this._url, idleTimeoutMillis: 30000});
        }
        return this._pool;
    }

    public connectClient(): Promise<PoolClient> {
        if (!this._pool) {
            this._pool = new Pool({max: 20, connectionString: this._url, idleTimeoutMillis: 30000});
        }
        return this._pool.connect();
    }

    public async connect() {
        console.log('Start connect to db:', this._hideUrl);
        this._pool.connect(function (err, client, done) {
            if (err) {
                console.error('Failed to connect to mongo on startup', err);
                throw err;
            }
            console.log('Connect to Postgres successful...');
        });
    }
}
