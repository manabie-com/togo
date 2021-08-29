import redis from "async-redis";
import objectHash from 'object-hash';
import _ from "lodash";
import { ClientOpts } from "redis";
import { injectable } from "inversify";

@injectable()
export class CacheClient {

    private readonly host: string = String(process.env.CACHE_HOST) || '127.0.0.1';
    private readonly port: number = Number(process.env.CACHE_PORT) || 6378;
    private readonly pass: string = String(process.env.CACHE_PASS) || '';

    public client: any;

    constructor() {
        const options: ClientOpts = {port: this.port, host: this.host};
        if (!!this.pass) {
            options.password = this.pass;
        }

        this.client = redis.createClient(options);
        this.client.on("error", function (err) {
            console.log("RedisClient error " + err);
        });
    }

    public static instance() {
        return new CacheClient();
    }

    public parseKey(key: string | object): string {
        if (_.isObject(key)) {
            return objectHash(key, {unorderedArrays: true});
        }
        return key;
    }

    public set(key: string | object, time: number, value: any) {
        const cacheKey: string = this.parseKey(key);
        const cacheValue: string = JSON.stringify({value: value});
        this.client.setex(cacheKey, time, cacheValue);
    }

    public get(key: string | object): any {
        const cacheKey: string = this.parseKey(key);
        const cacheValue: any = this.client.get(cacheKey);
        return cacheValue ? cacheValue.value : null;
    }
}
