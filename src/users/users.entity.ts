import {
  Entity,
  Column,
  PrimaryColumn,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';
@Entity()
export class User {
  @PrimaryColumn()
  username: string;

  @Column('text')
  password: string;

  @Column('int')
  limitPerDay: number;

  @CreateDateColumn()
  createAt: Date;

  @UpdateDateColumn()
  updateAt: Date;
}
