import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty } from 'class-validator';

export class UpsertUserDTO {
  @ApiProperty()
  @IsNotEmpty()
  name: string;

  @ApiProperty({ nullable: true, default: 0 })
  limitTaskInDay: number;
}
