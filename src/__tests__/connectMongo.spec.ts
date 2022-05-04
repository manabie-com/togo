import { MongoMemoryServer } from 'mongodb-memory-server';
import { connectMongo } from '../connectMongo';

describe('dbConnect', () => {
  describe('connectMongo', () => {
    it('should resolve if MONGO_URI is valid', async () => {
      const mongod = new MongoMemoryServer();
      const mongoDbUri = await mongod.getConnectionString();
      const success = await connectMongo({
        dbUri: mongoDbUri,
        dbName: 'test_db'
      }).then(() => 'success');
      await mongod.stop();
      expect(success).toEqual('success');
    });
    it('should reject if no MONGO_URI provided', async () => {
      const error = await connectMongo({
        dbUri: '',
        dbName: 'test_db'
      }).catch(() => 'error');
      expect(error).toEqual('error');
    });

    it('should reject if not MONGO_DB_NAME is provided', async () => {
      const mongod = new MongoMemoryServer();
      const mongoDbUri = await mongod.getConnectionString();
      const error = await connectMongo({
        dbUri: mongoDbUri,
        dbName: ''
      }).catch(() => 'error');
      await mongod.stop();
      expect(error).toEqual('error');
    });
  });
});
