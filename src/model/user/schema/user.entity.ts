import { Table, Column, Model, DataType, HasMany } from 'sequelize-typescript';
import { USER } from '../../../constance/variable';
import { LimitTask } from '../../limit-task/schema/limitTask.entity';
import { Task } from '../../task/schema/task.entity';

@Table({ modelName: USER, updatedAt: false, createdAt: false })
export class User extends Model {
  @Column({
    allowNull: false,
    primaryKey: true,
    type: DataType.INTEGER,
    autoIncrement: true,
  })
  id?: string;

  @Column({ allowNull: false, type: DataType.STRING })
  name: string;

  @Column({ allowNull: true, type: DataType.INTEGER })
  age?: number;

  // foreignKey
  @HasMany(() => Task)
  tasks: Task[];

  @HasMany(() => LimitTask)
  limitTasks: LimitTask[];
}
