import { ApiProperty } from '@nestjs/swagger';
import { IsEmpty, IsIn, IsNotEmpty, NotEquals, notEquals, ValidateIf } from "class-validator";
//import type { TaskEntity } from '../task.entity';

export class CreateUserDto {
  @IsNotEmpty()
  @ApiProperty()
  name: string;

  @ValidateIf(obj => { return obj >= 0})
  @ApiProperty()
  dailyTaskLimit: number;
}