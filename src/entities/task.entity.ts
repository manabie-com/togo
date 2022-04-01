import { ApiProperty } from '@nestjs/swagger';
import { Column, Entity, JoinColumn, ManyToOne, OneToMany } from 'typeorm';
import { TodoBase } from './todoBase.entity';
import { ToDoList } from './toDoList.entity';
import { User } from './user.entity';

export enum ETaskStatus {
  DO_TO = 'TO_DO',
  IN_PROGRESS = 'IN_PROGRESS',
  COMPLETE = 'COMPLETE',
}

@Entity()
export class Task extends TodoBase {
  constructor(parital: Partial<Task>) {
    super();
    Object.assign(this, parital);
  }
  @Column({ type: 'text', nullable: true })
  @ApiProperty()
  title: string;

  @Column({ type: 'text', nullable: true })
  @ApiProperty()
  desc: string;

  @Column({
    type: 'varchar',
    length: 30,
    nullable: false,
    default: ETaskStatus.DO_TO,
  })
  @ApiProperty({ enum: ETaskStatus })
  status: ETaskStatus;

  @Column({ type: 'timestamp', nullable: true })
  @ApiProperty()
  deadlineAt: Date;

  @ManyToOne(() => User)
  @JoinColumn()
  user?: User;

  @OneToMany(() => ToDoList, (toDoList) => toDoList.task)
  @JoinColumn()
  toDoList?: ToDoList[];
}
