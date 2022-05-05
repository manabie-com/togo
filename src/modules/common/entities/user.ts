import { Column, HasOne, Index, Model, Table } from 'sequelize-typescript';
import { UserSettingTask } from './user-setting-task';

@Table
export class User extends Model {
  @Column
  name: string;

  @Column
  @Index('user-name')
  username: string;

  @Column
  password: string;

  @Column({ field: 'created_at' })
  createdAt: Date;

  @Column({ field: 'updated_at' })
  updatedAt: Date;

  @HasOne(() => UserSettingTask)
  setting: UserSettingTask;
}
