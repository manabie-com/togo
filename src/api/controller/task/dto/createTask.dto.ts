import { Type } from "class-transformer";
import { IsNotEmpty, IsString } from "class-validator";

export class CreateTaskDto {
  @IsString()
  @IsNotEmpty()
  @Type(() => String)
  userId: string;

  @IsString()
  @IsNotEmpty()
  @Type(() => String)
  title: string;

  @IsString()
  @IsNotEmpty()
  @Type(() => String)
  desc: string;
}