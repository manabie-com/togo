import { Constants } from '../../../utils/constants';
import { Column, Entity, PrimaryGeneratedColumn, Unique } from 'typeorm';

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

    @Column()
    created_at: Date;

    @Column()
    updated_at: Date;
}
