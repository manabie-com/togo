import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';
import { StatusEnum, PriorityEnum } from './../../common/index';

@Entity('tasks') // here alias database name
export class Task {
  @PrimaryGeneratedColumn()
  id: number;

  // @Column({ nullable: true })
  // code: string;

  @Column({ nullable: true })
  title: string;

  @Column()
  assignee_id: number;

  @Column({ nullable: true })
  description: string;

  @Column({ nullable: true, default: PriorityEnum.Medium })
  priority: PriorityEnum;

  @Column({ default: StatusEnum.Todo })
  status: StatusEnum;

  @Column({ type: 'timestamp' })
  updated_time: Date;

  @Column({ type: 'timestamp' })
  created_time: Date;
}
