import sequelize, { Sequelize } from 'sequelize';
import {
  Table,
  Column,
  Model,
  DataType,
  BelongsTo,
  ForeignKey,
} from 'sequelize-typescript';
import { Optional } from 'sequelize/types';
import { TASK } from '../../../constance/variable';
import { User } from '../../user/schema/user.entity';

export interface TaskAttributes {
  id: number;
  title: string;
  desc: string;
  creationDate?: Date;
  userId: string;
}
export type TaskOptionalAttributes = 'id' | 'creationDate';

export type TaskCreationAttributes = Optional<
  TaskAttributes,
  TaskOptionalAttributes
>;

@Table({ modelName: TASK, updatedAt: false, createdAt: false })
export class Task extends Model<TaskAttributes, TaskCreationAttributes> {
  @Column({
    allowNull: false,
    primaryKey: true,
    type: DataType.INTEGER,
    autoIncrement: true,
  })
  id: string;

  @Column({ allowNull: false, type: DataType.STRING(100) })
  title: string;

  @Column({ allowNull: false, type: DataType.TEXT })
  desc: string;

  @Column({
    allowNull: false,
    type: DataType.DATE,
    defaultValue: Sequelize.fn('now'),
  })
  creationDate: Date;

  // foreignKey
  @ForeignKey(() => User)
  @Column({ allowNull: false, type: DataType.INTEGER })
  userId: string;

  @BelongsTo(() => User)
  user: User;
}
