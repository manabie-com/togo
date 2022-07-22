import { UserEntity } from "src/users/entities/user.entity";
import { BaseEntity, Column, Entity, JoinColumn, OneToOne, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class SettingEntity extends BaseEntity {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @OneToOne(() => UserEntity, { onDelete: 'CASCADE', onUpdate: 'CASCADE', nullable: false })
  @JoinColumn({ name: 'user_id' })
  user: UserEntity;

  @Column({ name: 'todo_per_day', default: 10 })
  todoPerday: number;

}
