import { Provider } from '@nestjs/common';
import { createConnection } from 'typeorm';
import { PostgresConnectionOptions } from 'typeorm/driver/postgres/PostgresConnectionOptions';

import { getConfig } from '../config';

const dbSettings = getConfig<IDBSettings>('DbSettings');

export const POSTGRES_CONNECTION = 'POSTGRE_CONNECTION';

export const DatabaseProvider: Provider[] = [
  {
    provide: POSTGRES_CONNECTION,
    useFactory: () =>
      createConnection({
        name: POSTGRES_CONNECTION,
        type: 'postgres',
        ...dbSettings,
        entities: [__dirname + '/entities/*.entity{.ts,.js}'],
        logging: false,
        synchronize: true
      } as PostgresConnectionOptions)
  }
];
