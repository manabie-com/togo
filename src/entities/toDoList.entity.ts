import { ApiProperty } from '@nestjs/swagger';
import { Column, Entity, JoinColumn, ManyToOne } from 'typeorm';
import { Task } from './task.entity';
import { TodoBase } from './todoBase.entity';

@Entity()
export class ToDoList extends TodoBase {
  constructor(parital: Partial<ToDoList>) {
    super();
    Object.assign(this, parital);
  }
  @Column('varchar')
  @ApiProperty()
  title: string;

  @Column('varchar')
  @ApiProperty()
  desc: string;

  @Column('boolean')
  @ApiProperty()
  isDone: boolean;

  @ManyToOne(() => Task)
  @JoinColumn()
  task: Task;
}
