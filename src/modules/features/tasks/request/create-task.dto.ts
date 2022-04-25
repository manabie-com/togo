import { IsArray, IsDefined, IsOptional, IsString } from 'class-validator';

export class CreateMultiTaskRequest {
  @IsDefined()
  @IsArray({ each: true })
  tasks: CreateTaskRequest[];
}

class CreateTaskRequest {
  @IsDefined()
  @IsString()
  title: string;

  @IsOptional()
  @IsString()
  description: string;

  @IsOptional()
  @IsString()
  note: string;

  @IsOptional()
  @IsArray()
  watchers: string[];

  @IsOptional()
  @IsArray()
  excutors: string[];
}
