import * as users from './User';
import { UserTaskConfig } from '@modules/users/entities/user-task-config.entity';

const tasks: UserTaskConfig[] = [];

for (const u of users) {
  tasks.push(
    new UserTaskConfig({
      userId: u.id,
      numberOfTaskPerDay: 5,
      date: new Date(),
    }),
  );
}

export = tasks;
