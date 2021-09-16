import { Global, HttpModule, Module } from '@nestjs/common';

import { DatabaseProvider, POSTGRES_CONNECTION } from './database.provider';
import { repositories, RepositoriesProvider } from './repositories.provider';

const providers = [...DatabaseProvider, ...RepositoriesProvider];

@Global()
@Module({
  providers,
  imports: [HttpModule],
  exports: [...providers, ...repositories, HttpModule, POSTGRES_CONNECTION]
})
export class SharedModule {}
