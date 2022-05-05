import {
  BelongsTo,
  Column,
  DataType,
  ForeignKey,
  Model,
  Table,
} from 'sequelize-typescript';
import { User } from './user';

@Table
export class UserSettingTask extends Model {
  @Column({
    field: 'user_id',
    type: DataType.INTEGER,
  })
  @ForeignKey(() => User)
  userId: number;

  @Column({ field: 'maximum_task' })
  maximum: number;

  @BelongsTo(() => User, {
    constraints: false,
    foreignKey: 'userId',
  })
  user: User;
  @Column({ field: 'created_at' })
  createdAt: Date;

  @Column({ field: 'updated_at' })
  updatedAt: Date;
}
