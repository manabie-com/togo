import { Module } from '@nestjs/common';
import { DatabaseModule } from 'src/database/database.module';
import { limitTaskProviders } from './provider/limitTask.provider';
import { LimitTaskService } from './service/limitTask.service';

@Module({
  imports: [DatabaseModule],
  providers: [
    LimitTaskService,
    ...limitTaskProviders,
  ],
  exports: [LimitTaskService]
})
export class LimitTaskModule {}