import { Column, PrimaryGeneratedColumn } from 'typeorm';

export interface IBaseEntity {
  id?: number;
  createdBy?: number;
  createdAt?: number;
  updatedBy?: number;
  updatedAt?: number;
}

export abstract class BaseEntity {
  protected constructor(props?: IBaseEntity) {
    Object.assign(this, props || {});
  }

  @PrimaryGeneratedColumn()
  id: number;

  @Column({ nullable: true, type: 'bigint' })
  createdBy?: number;

  @Column({ nullable: true, type: 'bigint' })
  createdAt?: number;

  @Column({ nullable: true, type: 'bigint' })
  updatedBy?: number;

  @Column({ nullable: true, type: 'bigint' })
  updatedAt?: number;
}
