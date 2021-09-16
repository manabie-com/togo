import { Column, Entity, ManyToOne } from 'typeorm';

import { BaseEntity, IBaseEntity } from './base.entity';
import { User } from './user.entity';

interface ITask extends IBaseEntity {
  text?: string;
  user?: User;
  isCompleted?: boolean;
}

@Entity('Task')
export class Task extends BaseEntity {
  constructor(props: ITask) {
    const { text, isCompleted, user, ...superItem } = props || {};

    super(superItem);

    Object.assign(this, text, user, isCompleted);
  }

  @Column({ type: 'varchar', width: 1024, nullable: true })
  text: string;

  @Column({ default: false })
  isCompleted?: boolean;

  @ManyToOne(() => User, (user) => user.tasks)
  user?: User;
}
