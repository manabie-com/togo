import { createClient, RedisClientType } from 'redis';
import { REDIS_URL } from '../config';
import logger from '../logger';

interface IRedisOptions {
  url: string;
}

class RedisService {
  private _client: RedisClientType;
  constructor(options: IRedisOptions) {
    this._client = createClient(options);
    this._init();
  }

  private _init = async (): Promise<void> => {
    await this._client.connect();
    logger.info(`Redis >>>> connected`);
  };

  set = async (key: string, value: string): Promise<void> => {
    await this._client.set(key, value);
  };

  get = async (key: string): Promise<string | null> => {
    return await this._client.get(key);
  };
}

export default new RedisService({
  url: REDIS_URL || ''
});
