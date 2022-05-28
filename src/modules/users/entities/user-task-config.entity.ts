import { ApiProperty } from '@nestjs/swagger';
import { Index, Entity, Column, JoinColumn, ManyToOne } from 'typeorm';
import { Expose, Exclude } from 'class-transformer';
import { CustomBaseEntity } from '@modules/common/entities/base.entity';
import { User } from './user.entity';

@Entity('user_task_configs')
export class UserTaskConfig extends CustomBaseEntity {
  static entityName: string = 'UserTaskConfig';

  @Expose()
  @ApiProperty()
  @Column({ name: 'number_of_task_per_day' })
  @Index('IDX_User_Task_Number_Of_Task_Per_Day')
  numberOfTaskPerDay: number;

  @Exclude()
  @Index('IDX_User_Task_User_Id')
  @Column({ name: 'user_id' })
  userId: string;

  @Expose()
  @ApiProperty()
  @ManyToOne(() => User)
  @JoinColumn({ name: 'user_id', referencedColumnName: 'id' })
  user: User;

  @Expose()
  @ApiProperty()
  @Column({ type: 'date' })
  date: Date;

  constructor(partial: Partial<UserTaskConfig>) {
    super();
    this.entityName = UserTaskConfig.entityName;
    Object.assign(this, partial);
  }
}
