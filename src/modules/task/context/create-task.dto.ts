import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsInt, IsNumber } from 'class-validator';
import { PriorityEnum, StatusEnum } from './../../../common/index';

export class CreateTaskDTO {
  @ApiProperty()
  @IsNotEmpty()
  title: string;

  @ApiProperty()
  @IsNotEmpty()
  @IsNumber()
  assignee_id: number;

  @ApiProperty({ nullable: true })
  description: string;

  @ApiProperty({ nullable: true, default: PriorityEnum.Medium })
  priority: PriorityEnum;

  @ApiProperty({ default: StatusEnum.Todo })
  status: StatusEnum;
}
