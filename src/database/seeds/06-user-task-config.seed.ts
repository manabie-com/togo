import { Factory, Seeder } from 'typeorm-seeding';
import { Connection } from 'typeorm';

import { cleanTable } from '@modules/common/utils/query';
import { UserTaskConfig } from '@modules/users/entities/user-task-config.entity';
import { User } from '@modules/users/entities/user.entity';
import { Role } from '@modules/roles/entities/role.entity';
import { USER_ROLE_NAME } from '@modules/common/constant';

export default class CreateUserTaskConfigSeeder implements Seeder {
  tableName = 'user_task_configs';

  public async run(factory: Factory, connection: Connection): Promise<any> {
    await cleanTable(connection, this.tableName);

    const userRole = await Role.findOne({ where: { name: USER_ROLE_NAME } });

    const users = await connection.getRepository(User).find({ where: { roleId: userRole.id } });

    const userTaskConfigs: UserTaskConfig[] = users.map(
      (user) => new UserTaskConfig({ userId: user.id, numberOfTaskPerDay: 5, date: new Date() }),
    );

    await connection.createQueryBuilder().insert().into(this.tableName).values(userTaskConfigs).execute();
  }
}
