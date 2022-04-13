import { Column, Entity } from 'typeorm';
import { BaseEntity } from '../../../common/entity/base.entity';

@Entity({ name: 'user' })
export class UserEntity extends BaseEntity {
  @Column({
    name: 'user_name',
    type: 'varchar',
  })
  username: string;

  @Column({
    name: 'password',
    type: 'varchar',
  })
  password: string;

  @Column({
    name: 'max_task',
    type: 'int',
    default: 3,
  })
  maxTask: number;
}
