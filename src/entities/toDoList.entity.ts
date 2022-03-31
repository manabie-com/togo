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
  title: string;

  @Column('varchar')
  desc: string;

  @Column('boolean')
  isDone: boolean;

  @ManyToOne(() => Task)
  @JoinColumn()
  task: Task;
}
