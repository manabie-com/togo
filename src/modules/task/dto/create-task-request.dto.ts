import { IsNotEmpty, IsOptional, IsString } from 'class-validator';

export class CreateTaskRequestDto {
  @IsString()
  @IsNotEmpty()
  title: string;

  @IsString()
  @IsOptional()
  content: string;
}
