import { Column, Entity, OneToMany } from 'typeorm';

import { BaseEntity, IBaseEntity } from './base.entity';
import { Task } from './task.entity';

interface IUser extends IBaseEntity {
  name?: string;
  hash?: string;
  salt?: string;
}

@Entity('User')
export class User extends BaseEntity {
  constructor(props: IUser) {
    const { name, hash, salt, ...superItem } = props || {};

    super(superItem);

    Object.assign(this, name, hash, salt);
  }

  @Column({ type: 'varchar', width: 64, nullable: true })
  name: string;

  @Column({ type: 'varchar', width: 64, nullable: true })
  email: string;

  @Column({ type: 'varchar', nullable: false })
  hash: string;

  @Column({ type: 'varchar', width: 64, nullable: false })
  salt: string;

  @OneToMany(() => Task, (task) => task.user)
  tasks: Task[];
}
