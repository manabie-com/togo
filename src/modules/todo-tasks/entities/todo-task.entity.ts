import { ApiProperty } from '@nestjs/swagger';
import { Column, Entity, Index, JoinColumn, ManyToOne } from 'typeorm';

import { CustomBaseEntity } from '@modules/common/entities/base.entity';
import { User } from '@modules/users/entities/user.entity';
import { TaskStatusEnum } from '../enums/task-status.enum';
import { Exclude } from 'class-transformer';

@Entity('tasks')
export class TodoTask extends CustomBaseEntity {
  static entityName: string = 'TodoTask';

  @ApiProperty({ type: 'string', example: 'TODO-1000: [BUG][Urgent] User cannot log in' })
  @Column({ type: 'nvarchar', length: 200 })
  summary: string;

  @ApiProperty({ type: 'string', example: 'User cannot log in the system with username and password' })
  @Column({ nullable: true, type: 'nvarchar', length: 500 })
  description: string;

  @Exclude()
  @Index('IDX_TodoTask_Assignee_Id')
  @Column({ name: 'assignee_id', nullable: true })
  assigneeId: string;

  // 1 User has Many Tasks
  @ManyToOne(() => User)
  @JoinColumn({ name: 'assignee_id', referencedColumnName: 'id' })
  assignee: User;

  @ApiProperty({
    type: TaskStatusEnum,
    enum: TaskStatusEnum,
    enumName: 'TaskStatusEnum',
    example: TaskStatusEnum.Todo,
  })
  @Column({ default: TaskStatusEnum.Todo })
  status: TaskStatusEnum;

  constructor(partial?: Partial<TodoTask>) {
    super();
    this.entityName = TodoTask.entityName;
    Object.assign(this, partial);
  }
}
