import {
  CreateDateColumn,
  PrimaryGeneratedColumn,
  UpdateDateColumn,
} from 'typeorm';

export class TodoBase {
  @PrimaryGeneratedColumn()
  id: number;

  @UpdateDateColumn()
  createdAt: Date;

  @CreateDateColumn()
  deletedAt: Date;
}
