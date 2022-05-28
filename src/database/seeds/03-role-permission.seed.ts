import { Factory, Seeder } from 'typeorm-seeding';
import { Connection } from 'typeorm';
import { rolePermissions } from './data';
import { cleanTable } from '@modules/common/utils/query';

export default class CreateRolePermissionSeeder implements Seeder {
  tableName = 'role_permissions';

  public async run(factory: Factory, connection: Connection): Promise<any> {
    await cleanTable(connection, this.tableName);
    const data = rolePermissions.map((x) => {
      return {
        role_id: x.roleId,
        permission_id: x.permissionId,
      };
    });

    await connection.createQueryBuilder().insert().into(this.tableName).values(data).execute();
  }
}
