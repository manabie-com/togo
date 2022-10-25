import {
  Entity,
  Column,
  PrimaryColumn,
  OneToMany,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';
import { Todo } from '../todo/todo.entity';

@Entity()
export class User {
  @PrimaryColumn()
  username: string;

  @Column('text')
  password: string;

  @Column('int')
  limitPerDay: number;

  @OneToMany('Todo', 'user', { cascade: true })
  todos: Todo[];

  @CreateDateColumn()
  createAt: Date;

  @UpdateDateColumn()
  updateAt: Date;
}
