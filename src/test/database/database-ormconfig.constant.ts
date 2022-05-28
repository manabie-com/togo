import { Permission } from '@modules/permissions/permission.entity';
import { Role } from '@modules/roles/entities/role.entity';
import { TodoTask } from '@modules/todo-tasks/entities/todo-task.entity';
import { UserTaskConfig } from '@modules/users/entities/user-task-config.entity';
import { User } from '@modules/users/entities/user.entity';

export const getOrmConfig = (): any => {
  const ormConfig = {
    type: 'sqlite',
    database: ':memory:',
    entities: [User, Role, Permission, TodoTask, UserTaskConfig],
    synchronize: true,
    dropSchema: true,
    timezone: 'utc',
    keepConnectionAlive: true,
  };

  return ormConfig;
};
