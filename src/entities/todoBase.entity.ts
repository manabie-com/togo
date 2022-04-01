import { ApiProperty } from '@nestjs/swagger';
import {
  CreateDateColumn,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';

export class TodoBase {
  @PrimaryGeneratedColumn()
  @ApiProperty()
  id: number;

  @UpdateDateColumn()
  @ApiProperty()
  createdAt: Date;

  @CreateDateColumn()
  @ApiProperty()
  deletedAt: Date;
}
