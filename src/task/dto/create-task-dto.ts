import { ApiProperty } from '@nestjs/swagger';
import { IsEmpty, IsIn, IsNotEmpty, NotEquals, notEquals, ValidateIf } from "class-validator";
//import type { TaskEntity } from '../task.entity';

export class CreateTaskDto {
  @IsNotEmpty()
  @ApiProperty()
  title: string;

  @IsNotEmpty()
  @ApiProperty()
  content: string;

  @IsNotEmpty()
  @ApiProperty()
  userId: string;

  @ApiProperty()
  dateTime: Date;
  // constructor(task: TaskEntity) {
  //   this.title = task.title;
  //   this.content = task.content;
  //   this.dateTime = task.dateTime;
  //   this.createdAt = task.createdAt;
  //   this.updatedAt = task.updatedAt;
  // }
}