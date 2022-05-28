import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsUUID } from 'class-validator';

export class AssignTaskDto {
  @ApiProperty({ example: '6780433E-F36B-1410-86BA-002BD86783D2' })
  @IsNotEmpty()
  @IsUUID()
  taskId: string;
}
