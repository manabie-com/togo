import {
  Entity,
  BaseEntity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  OneToMany,
  Index,
  DeleteDateColumn,
} from 'typeorm';
import { Task } from './task.entity';

@Entity()
export class User extends BaseEntity {
  @PrimaryGeneratedColumn('increment')
  id: number;

  @Index({ unique: true })
  @Column({ type: 'varchar' })
  userId: string;

  @Column({ type: 'varchar', select: false })
  password: string;

  @Column({ type: 'int4', default: 5 })
  maxTodo: number;

  @OneToMany(() => Task, (task) => task.user)
  tasks: Task[];

  @CreateDateColumn({ type: 'timestamptz' })
  createdAt: Date;

  @UpdateDateColumn({ type: 'timestamptz', nullable: true })
  updatedAt: Date;
}
