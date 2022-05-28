import { ApiProperty } from '@nestjs/swagger';
import { Expose } from 'class-transformer';

export class JWTPayload {
  @Expose()
  @ApiProperty()
  readonly userId: string;

  @Expose()
  @ApiProperty()
  readonly username: string;

  @Expose()
  @ApiProperty()
  readonly role?: string;

  constructor(payload: { userId: string; username: string; role?: string }) {
    this.userId = payload.userId;
    this.username = payload.username;
    this.role = payload.role;
  }
}
