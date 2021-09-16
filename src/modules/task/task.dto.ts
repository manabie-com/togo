import { ApiProperty } from '@nestjs/swagger';
import { IsBoolean, IsNotEmpty, IsOptional, IsString } from 'class-validator';

export class CreateTaskDto {
  @ApiProperty({ name: 'text', example: 'Listen to music', required: true })
  @IsString()
  @IsNotEmpty()
  readonly text: string;
}

export class UpdateTaskDto {
  @ApiProperty({ name: 'text', example: 'Listen to music', required: false })
  @IsString()
  @IsOptional()
  readonly text: string;

  @ApiProperty({ name: 'isCompleted', example: true, required: false })
  @IsBoolean()
  @IsOptional()
  readonly isCompleted: boolean;
}
