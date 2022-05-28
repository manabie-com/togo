import { Factory, Seeder } from 'typeorm-seeding';
import { Connection } from 'typeorm';
import { createForeignKeys } from '@modules/common/utils/query';

export default class CreateForeignKeysSeeder implements Seeder {
  public async run(factory: Factory, connection: Connection): Promise<any> {
    await createForeignKeys(connection);
  }
}
