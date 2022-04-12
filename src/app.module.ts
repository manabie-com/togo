import { Module } from '@nestjs/common';
import { PgModule } from './common/database/pg.module';
import { EnvironmentModule } from './common/environment/environment.module';

@Module({
  imports: [
    EnvironmentModule,
    PgModule.register({
      entities: [],
    }),
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
