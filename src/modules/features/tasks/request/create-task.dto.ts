import { ApiProperty } from '@nestjs/swagger';
import { IsArray, IsDefined, IsOptional, IsString } from 'class-validator';

class CreateTaskRequest {
  @IsDefined()
  @IsString()
  @ApiProperty()
  title: string;

  @IsOptional()
  @IsString()
  @ApiProperty()
  description: string;

  @IsOptional()
  @IsString()
  @ApiProperty()
  note: string;

  @IsOptional()
  @IsArray()
  @ApiProperty({ type: [String] })
  watchers: string[];

  @IsOptional()
  @IsArray()
  @ApiProperty({ type: [String] })
  excutors: string[];
}

export class CreateMultiTaskRequest {
  @IsDefined()
  @ApiProperty({ type: [CreateTaskRequest] })
  tasks: CreateTaskRequest[];
}
