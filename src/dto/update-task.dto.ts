import { ApiProperty } from '@nestjs/swagger';
import { IsEnum, IsNumber, Min, ValidateIf } from 'class-validator';
import { isNil } from 'lodash';
import { ETaskStatus } from 'src/entities/task.entity';
import { CreateTaskDto } from './create-task.dto';

export class UpdateTaskDto extends CreateTaskDto {
  @ApiProperty({ enum: ETaskStatus })
  @ValidateIf((obj, value) => !isNil(value))
  @IsEnum(ETaskStatus)
  status: ETaskStatus;

  @ApiProperty()
  @ValidateIf((obj, value) => !isNil(value))
  @IsNumber()
  @Min(1)
  userId: number;
}
