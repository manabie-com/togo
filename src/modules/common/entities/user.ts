import { Column, Model, Table } from 'sequelize-typescript';

@Table
export class User extends Model {
  @Column
  name: string;

  @Column
  username: string;

  @Column
  password: string;

  @Column
  createdAt: Date;

  @Column
  updatedAt: Date;
}
