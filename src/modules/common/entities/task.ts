import { Column, Model, Table } from 'sequelize-typescript';

@Table
export class Task extends Model {
  @Column
  title: string;

  @Column
  description: string;

  @Column
  note: string;

  @Column
  status: TaskStatus;

  @Column
  createdBy: string;

  @Column
  createdAt: Date;

  @Column
  updatedAt: Date;
}
