import { ApiProperty } from '@nestjs/swagger';
import { Expose } from 'class-transformer';

export class Token {
  @Expose({ name: 'access_token' })
  @ApiProperty({ name: 'access_token' })
  readonly accessToken: string;

  @Expose({ name: 'refresh_token' })
  @ApiProperty({ name: 'refresh_token' })
  readonly refreshToken: string;

  @Expose({ name: 'token_type' })
  @ApiProperty({ name: 'token_type' })
  readonly tokenType: string;

  @Expose({ name: 'expires_in' })
  @ApiProperty({ name: 'expires_in' })
  readonly expiresIn: number;

  constructor(accessToken: string, refreshToken: string, expiresIn: number) {
    this.expiresIn = expiresIn;
    this.accessToken = accessToken;
    this.refreshToken = refreshToken;
    this.tokenType = 'Bearer';
  }
}
