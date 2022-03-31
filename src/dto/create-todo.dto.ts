import { IsNotEmpty, IsString } from 'class-validator';

export class CreateToDoDto {
  @IsNotEmpty()
  @IsString()
  title: string;

  @IsNotEmpty()
  @IsString()
  desc: string;
}
