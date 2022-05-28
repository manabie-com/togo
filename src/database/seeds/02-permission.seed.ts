import { Factory, Seeder } from 'typeorm-seeding';
import { Connection } from 'typeorm';
import { permissions } from './data';
import { cleanTable } from '@modules/common/utils/query';

export default class CreatePermissionSeeder implements Seeder {
  tableName = 'permissions';

  public async run(factory: Factory, connection: Connection): Promise<any> {
    await cleanTable(connection, this.tableName);

    const chunk = 100;

    const data = permissions.map((x) => {
      return {
        id: x.id,
        name: x.name,
        resource: x.resource,
        action: x.action,
      };
    });

    const { length } = data;

    for (let i = 0, j = length; i < j; i += chunk) {
      const permissionChunk = data.slice(i, i + chunk);

      await connection.createQueryBuilder().insert().into(this.tableName).values(permissionChunk).execute();
    }
  }
}
