import { UserEntity } from "src/users/entities/user.entity";
import { BaseEntity, Column, Entity, JoinColumn, ManyToOne, PrimaryGeneratedColumn } from "typeorm";

@Entity()
export class TodoEntity extends BaseEntity {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @ManyToOne(() => UserEntity, { onDelete: 'CASCADE', onUpdate: 'CASCADE' })
  @JoinColumn()
  user: UserEntity;

  @Column()
  title: string;

  @Column({ type: 'date' })
  date: Date;

}
