import { Body, Controller, Get, Param, Patch, Put } from '@nestjs/common';
import { ApiTags } from '@nestjs/swagger';
import { UpdateSettingDto } from './dto/update-setting.dto';
import { SettingsService } from './settings.service';

@ApiTags('Settings')
@Controller('users/:userId/settings')
export class SettingsController {
  constructor(private readonly settingsService: SettingsService) { }

  @Get('')
  async findOne(@Param('userId') userId: string) {
    return this.settingsService.findByUserId(userId);
  }

  @Put('')
  async update(@Param('userId') userId: string, @Body() updateSettingDto: UpdateSettingDto) {
    const setting = await this.settingsService.findByUserId(userId);
    return this.settingsService.update(setting.id, updateSettingDto);
  }

}
