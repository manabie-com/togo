import { Module } from '@nestjs/common';
import { EnvironmentModule } from '../environment/environment.module';
import { PgService } from './services/pg.service';

@Module({
  imports: [EnvironmentModule],
  providers: [PgService],
  exports: [PgService],
})
export class PgSubModule {}
