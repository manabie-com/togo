import { connect, connection } from 'mongoose';
import logger from './logger';

export interface IConnectMongo {
  dbUri?: string;
  user?: string;
  pass?: string;
  dbName?: string;
}

export const connectMongo = ({
  dbUri = '',
  user = '',
  pass = '',
  dbName = ''
}: IConnectMongo) =>
  new Promise<void>((resolve, reject) => {
    if (!dbUri) {
      logger.error(`No MONGO_URI is provided`);
      return reject(`No MONGO_URI is provided`);
    }

    if (!dbName) {
      logger.error(`No MONGO_DB_NAME is provided`);
      return reject(`No MONGO_DB_NAME is provided`);
    }

    connection.once('open', () => {
      resolve();
    });

    connection.on('error', (err: any) => {
      logger.error('error while connecting to mongodb', { message: err.stack });
      reject(err);
    });

    connect(dbUri, { user, pass, dbName });
  });
