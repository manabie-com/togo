import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsString } from 'class-validator';

export class LoginPayload {
  @ApiProperty()
  @IsDefined()
  @IsString()
  readonly username: string;

  @ApiProperty()
  @IsDefined()
  @IsString()
  readonly password: string;
}
