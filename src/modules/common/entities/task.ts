import sequelize from 'sequelize';
import { Column, Model, Table } from 'sequelize-typescript';

@Table
export class Task extends Model {
  @Column
  title: string;

  @Column
  description: string;

  @Column
  note: string;

  @Column({
    type: sequelize.DataTypes.ENUM,
    values: ['TO_DO', 'IN_PROGESS', 'REVIEW', 'RE_OPEN', 'TESTING', 'DONE'],
  })
  status: TaskStatus;

  @Column({ field: 'created_by' })
  createdBy: number;

  @Column({ field: 'created_at' })
  createdAt: Date;

  @Column({ field: 'updated_at' })
  updatedAt: Date;
}
