
import { Column, Model, Table } from 'sequelize-typescript';

@Table
export class UserSettingTask extends Model{

  @Column
  userId: string;

  @Column
  perTaskOnDate: number;
}

