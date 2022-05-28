import { ApiProperty } from '@nestjs/swagger';
import { Optional } from '@nestjs/common';
import { Index, Entity, Column, JoinColumn, ManyToOne, OneToMany } from 'typeorm';
import { Expose, Exclude } from 'class-transformer';
import { IsString, IsEmail, MaxLength } from 'class-validator';

import { CustomBaseEntity } from '@modules/common/entities/base.entity';
import { Role } from '@modules/roles/entities/role.entity';
import { UserTaskConfig } from './user-task-config.entity';
import { formatDate } from '@modules/common/utils/date-time.helper';
import { TodoTask } from '@modules/todo-tasks/entities/todo-task.entity';

@Entity('users')
export class User extends CustomBaseEntity {
  static entityName: string = 'User';

  @Exclude()
  @ApiProperty()
  @Column({ length: 100 })
  @Index('UIDX_User_Username', { unique: true })
  @MaxLength(100)
  @IsString()
  username: string;

  @Expose()
  @ApiProperty()
  @Column({ length: 100, nullable: true })
  @Index('UIDX_User_Email', { unique: true, where: 'email IS NOT NULL' })
  @MaxLength(100)
  @IsEmail()
  email?: string;

  @Expose()
  @ApiProperty()
  @MaxLength(100)
  @Optional()
  @Column({ length: 100, nullable: true })
  @Index('IDX_User_Mobile')
  mobile?: string;

  @Expose()
  @ApiProperty()
  @MaxLength(100)
  @Index('IDX_User_Display_Name')
  @Column({ length: 512, name: 'display_name', nullable: true })
  displayName?: string;

  @Exclude()
  @Column({ length: 100, nullable: true })
  password?: string;

  @Expose()
  @ApiProperty()
  get roleName(): string {
    return this.role?.name || '';
  }

  @Exclude()
  @Index('IDX_User_Role_Id')
  @Column({ name: 'role_id', nullable: true })
  roleId?: string;

  @Expose()
  @ApiProperty()
  @ManyToOne(() => Role)
  @JoinColumn({ name: 'role_id', referencedColumnName: 'id' })
  role: Role;

  @Expose()
  @ApiProperty()
  @OneToMany(() => UserTaskConfig, (task) => task.user)
  @JoinColumn({ name: 'user_id', referencedColumnName: 'id' })
  userTaskConfigs: UserTaskConfig[];

  get numberOfTaskToday() {
    const numberOfTask = this.userTaskConfigs.find((x) => formatDate(new Date(), 'yyyy-MM-dd') === x.date.toString());

    return numberOfTask?.numberOfTaskPerDay || 1;
  }

  constructor(partial: Partial<User>) {
    super();
    this.entityName = User.entityName;
    Object.assign(this, partial);
  }

  async myTotalTask(): Promise<number> {
    return await TodoTask.count({ where: { assigneeId: this.id } });
  }

  async canPickMoreTask(): Promise<boolean> {
    return (await this.myTotalTask()) < this.numberOfTaskToday;
  }
}
