import { ApiProperty } from '@nestjs/swagger';
import { Expose } from 'class-transformer';

export class MessageResult {
  @Expose()
  @ApiProperty({ example: ['messageKeyHere'] })
  readonly message: string | string[];
}
