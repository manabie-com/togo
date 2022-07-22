import { OmitType, PartialType } from '@nestjs/swagger';
import { SettingDto } from './setting.dto';

export class UpdateSettingDto extends PartialType(OmitType(SettingDto, ['id', 'user'])) { }
