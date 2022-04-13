import { IsString } from 'class-validator';

export class UserAuthDto {
  @IsString()
  username: string;

  @IsString()
  password: string;
}
