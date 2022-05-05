import { Module } from '@nestjs/common';
import { SequelizeModule } from '@nestjs/sequelize';
import { User, UserSettingTask } from 'src/modules/common/entities';
import { InitDataService } from './init-data.service';

@Module({
  imports: [SequelizeModule.forFeature([User, UserSettingTask])],
  providers: [InitDataService],
})
export class InitDataModule {}
