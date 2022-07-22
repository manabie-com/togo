import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { SettingEntity } from './entities/setting.entity';
import { SettingsController } from './settings.controller';
import { SettingsService } from './settings.service';

export const entities = [SettingEntity];

@Module({
  imports: [TypeOrmModule.forFeature(entities)],
  controllers: [SettingsController],
  providers: [SettingsService]
})
export class SettingsModule {}
