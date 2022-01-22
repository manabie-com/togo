import { Constants } from '../../../utils/constants';
import { Column, Entity, OneToMany, PrimaryGeneratedColumn, Unique } from 'typeorm';
import { TaskEntity } from './task.entity';
import { Min } from 'class-validator';

@Entity({ name: 'users' })
@Unique("unq_user_email", ["email"])
export class UserEntity {
    constructor(partial?: Partial<UserEntity>) {
        Object.assign(this, partial);
    }

    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    email: string;

    @Column()
    password: string;

    @Column()
    fullName: string;

    @Column({
        default: Constants.USER_ROLE
    })
    role: string;

    @OneToMany(type => TaskEntity, task => task.owner)
    listTask: TaskEntity[]

    @Column({
        unsigned: true,
        default: Constants.MAX_TASK_PER_DAY_DEFAULT
    })
    max_task_per_day: number;

    @Column()
    lasted_date_task: Date;

    @Column({
        unsigned: true,
        default: Constants.MAX_TASK_PER_DAY_DEFAULT
    })
    task_left: number;

    @Column()
    created_at: Date;

    @Column()
    updated_at: Date;
}
