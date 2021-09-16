import { ApiProperty } from '@nestjs/swagger';
import { IsEmail, IsNotEmpty, IsString } from 'class-validator';

export class UserRegisterDto {
  @ApiProperty({ name: 'name', example: 'Bob', required: true })
  @IsString()
  @IsNotEmpty()
  readonly name: string;

  @ApiProperty({ name: 'email', example: 'demo@demo.com', required: true })
  @IsString()
  @IsEmail()
  @IsNotEmpty()
  readonly email: string;

  @ApiProperty({ name: 'password', example: '12345@bC', required: true })
  @IsString()
  @IsNotEmpty()
  readonly password: string;
}

export class UserLoginDto {
  @ApiProperty({ name: 'email', example: 'demo@demo.com', required: true })
  @IsEmail()
  @IsNotEmpty()
  readonly email: string;

  @ApiProperty({ name: 'password', example: '12345@bC', required: true })
  @IsString()
  @IsNotEmpty()
  readonly password: string;
}
