import { Sequelize } from 'sequelize';
import {
  Table,
  Column,
  Model,
  DataType,
  BelongsTo,
  ForeignKey,
} from 'sequelize-typescript';
import { Optional } from 'sequelize/types';
import { LIMIT_TASK } from 'src/constance/variable';
import { User } from '../../user/schema/user.entity';

export interface LimitTaskAttributes {
  id: number
  limitNumber: number;
  creationDate?: Date;
  userId: string;
}
export type LimitTaskOptionalAttributes = "id" | "creationDate";

export type LimitTaskCreationAttributes = Optional<LimitTaskAttributes, LimitTaskOptionalAttributes>

@Table({ modelName: LIMIT_TASK, updatedAt: false, createdAt: false })
export class LimitTask extends Model<LimitTaskAttributes, LimitTaskCreationAttributes> {
  @Column({
    allowNull: false,
    primaryKey: true,
    type: DataType.INTEGER,
    autoIncrement: true,
  })
  id: string;

  @Column({ allowNull: false, type: DataType.INTEGER  })
  limitNumber: number;

  @Column({ allowNull: false, type: DataType.DATE, defaultValue: Sequelize.fn('now') })
  creationDate: Date;

  // foreignKey
  @ForeignKey(() => User)
  @Column({ allowNull: false, type: DataType.INTEGER })
  userId: string;

  @BelongsTo(() => User)
  user: User;
}