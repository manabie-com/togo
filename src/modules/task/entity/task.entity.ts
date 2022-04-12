import { Column, Entity } from 'typeorm';
import { BaseEntity } from '../../../common/entity/base.entity';

@Entity({ name: 'task' })
export class TaskEntity extends BaseEntity {
  @Column({
    name: 'title',
    type: 'varchar',
  })
  title: string;

  @Column({
    name: 'content',
    type: 'varchar',
  })
  content: string;

  @Column({
    name: 'user_id',
    type: 'uuid',
  })
  userId: string;
}
