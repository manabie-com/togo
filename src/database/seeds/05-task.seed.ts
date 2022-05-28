import { Factory, Seeder } from 'typeorm-seeding';
import { faker } from '@faker-js/faker';
import { Connection } from 'typeorm';

import { cleanTable } from '@modules/common/utils/query';
import { TodoTask } from '@modules/todo-tasks/entities/todo-task.entity';

const tasks: TodoTask[] = [];

for (let i = 1; i <= 20; i++) {
  tasks.push(
    new TodoTask({
      summary: `TODO-${i}: Buy ${faker.commerce.productName()}`,
      description: faker.lorem.paragraph(),
    }),
  );
}

export default class CreateTaskSeeder implements Seeder {
  tableName = 'tasks';

  public async run(factory: Factory, connection: Connection): Promise<any> {
    await cleanTable(connection, this.tableName);

    await connection.createQueryBuilder().insert().into(this.tableName).values(tasks).execute();
  }
}
