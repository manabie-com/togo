import { ApiProperty } from '@nestjs/swagger';
import { IsBoolean, ValidateIf } from 'class-validator';
import { isNil } from 'lodash';
import { CreateToDoDto } from './create-todo.dto';

export class UpdateToDoDto extends CreateToDoDto {
  @ApiProperty()
  @ValidateIf((obj, value) => !isNil(value))
  @IsBoolean()
  isDone: boolean;
}
