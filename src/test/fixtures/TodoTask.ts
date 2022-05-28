import { TodoTask } from '@modules/todo-tasks/entities/todo-task.entity';
import { faker } from '@faker-js/faker';
import * as users from '../../../src/test/fixtures/User';

const tasks: TodoTask[] = [];

const { length } = users;

for (let i = 1; i <= 10; i++) {
  tasks.push(
    new TodoTask({
      summary: `TODO-${i}: Buy ${faker.commerce.productName()}`,
      description: faker.lorem.paragraph(),
      assigneeId: users[i <= length ? i - 1 : 1].id,
    }),
  );
}

export = tasks;
