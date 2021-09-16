import { Provider } from '@nestjs/common';
import { Connection } from 'typeorm';

import { POSTGRES_CONNECTION } from './database.provider';
import { TaskRepository } from './repositories/task.repository';
import { UserRepository } from './repositories/user.repository';

const repositories = [UserRepository, TaskRepository];

const RepositoriesProvider: Provider[] = [];

for (const repository of repositories) {
  RepositoriesProvider.push({
    provide: repository,
    useFactory: (connection: Connection) =>
      connection.getCustomRepository(repository),
    inject: [POSTGRES_CONNECTION]
  });
}

export { RepositoriesProvider, repositories };
