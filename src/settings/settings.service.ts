import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { UpdateSettingDto } from './dto/update-setting.dto';
import { SettingEntity } from './entities/setting.entity';

@Injectable()
export class SettingsService {
  constructor(
    @InjectRepository(SettingEntity) private settingsRepository: Repository<SettingEntity>,
  ) { }

  async findByUserId(userId: string) {
    const setting = await this.settingsRepository.findOne({ where: { user: { id: userId } } });

    if (!setting) {
      throw new NotFoundException('Setting not found');
    }

    return setting;
  }

  async update(id: string, updateSettingDto: UpdateSettingDto) {
    const setting = await this.settingsRepository.findOne({ where: { id } });

    if (!setting) {
      throw new NotFoundException('Setting not found');
    }

    await this.settingsRepository.update(id, updateSettingDto);

    return this.settingsRepository.findOne({ where: { id } });
  }

}
