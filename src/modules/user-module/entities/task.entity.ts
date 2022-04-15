import { Column, Entity, ManyToOne, PrimaryGeneratedColumn, Unique } from 'typeorm';
import { UserEntity } from './user.entity';

@Entity({ name: 'tasks' })
export class TaskEntity {
    constructor(partial?: Partial<TaskEntity>) {
        Object.assign(this, partial);
    }

    @PrimaryGeneratedColumn()
    id: number;

    @Column()
    name: string;

    @Column()
    note: string;

    @ManyToOne(type => UserEntity, user => user.listTask)
    owner: UserEntity

    @Column()
    created_at: Date;
}
