import { Type } from 'class-transformer';
import { IsNotEmpty, IsString } from 'class-validator';

export class CreateTaskDto {
  @IsString()
  @Type(() => String)
  @IsNotEmpty()
  userId: string;

  @IsString()
  @Type(() => String)
  @IsNotEmpty()
  title: string;

  @IsString()
  @Type(() => String)
  @IsNotEmpty()
  desc: string;
}
