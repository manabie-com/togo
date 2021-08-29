import dotenv from "dotenv";
import mongoose from "mongoose";
import { config, down, up } from "migrate-mongo/lib/migrate-mongo";
import { database } from "migrate-mongo";
import { CONNECTION_OPTIONS, MIGRATE_CONFIG } from "../config";
import { setLogLevel } from '@typegoose/typegoose';
import * as logger from "loglevel";
import { injectable } from "inversify";

@injectable()
export class MongoClient {

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
        this._url = `mongodb://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}`;
        this._hideUrl = `mongodb://${DB_USER}:xxx@${DB_HOST}:${DB_PORT}/${DB_NAME}`;
        if (!DB_USER || !DB_PASS) {
            this._url = `mongodb://${DB_HOST}:${DB_PORT}/${DB_NAME}`;
            this._hideUrl = `mongodb://${DB_HOST}:${DB_PORT}/${DB_NAME}`;
        }
    }

    public static newInstance(): MongoClient {
        return new MongoClient();
    }

    public get url(): string {
        return this._url;
    }

    public async connect() {
        console.log('Start connect to db:', this._hideUrl);
        mongoose.connect(this._url, CONNECTION_OPTIONS).then(() => {
            console.log('Connect to MongoDB successful...');
            if (process.env.DB_DEBUG === 'true') {
                this.debug(mongoose);
            }
        }).catch((err: Error) => {
            console.error('Failed to connect to mongo on startup', err);
        });
    }

    public async migrate(action: 'up' | 'down') {
        MIGRATE_CONFIG.mongodb.url = this._url;
        config.set(MIGRATE_CONFIG);
        const {db, client} = await database.connect();

        if (action === 'up') {
            try {
                let result = await up(db, client);
                console.log('Migrating MongoDB database up successfully with result:', result);
            } catch (err) {
                console.error('Migrating MongoDB database up failure with error:', err);
            } finally {
                await client.close();
            }
            return;
        }

        try {
            let result = await down(db, client);
            console.log('Migrating MongoDB database down successfully with result:', result);
        } catch (err) {
            console.error('Migrating MongoDB database down failure with error:', err);
        } finally {
            await client.close();
        }
    }

    private debug(mongoose: any) {
        mongoose.set('debug', function (collection: any, method: any, query: any, doc: any) {
            if (method !== 'createIndex') {
                console.log('collection', collection);
                console.log('method', method);
                console.log('query', query);
                console.log('doc', doc);
                console.log('----------------------');
            }
        });
    }
}
