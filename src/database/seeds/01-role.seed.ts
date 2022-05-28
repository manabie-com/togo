import { Factory, Seeder } from 'typeorm-seeding';
import { Connection } from 'typeorm';
import { roles } from './data';
import { cleanTable } from '@modules/common/utils/query';

const tableName = 'roles';

export default class CreateRolSeeder implements Seeder {
  public async run(factory: Factory, connection: Connection): Promise<any> {
    await cleanTable(connection, tableName);

    const values = roles.map((item) => {
      return `('${item.id}', N'${item.name}')`;
    });

    await connection.query(`
        INSERT INTO "${tableName}" ("id", "name") VALUES ${values.join(',')}`);
  }
}
