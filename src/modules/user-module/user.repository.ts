import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Constants } from '../../utils/constants';
import { Repository } from 'typeorm';
import { RegisterDTO } from './dto/register.dto';
import { UserEntity } from './entities/user.entity';


@Injectable()
export class UserRepository {
    constructor(
        @InjectRepository(UserEntity)
        private readonly userRepository: Repository<UserEntity>,
    ) { }

    async register(data: RegisterDTO, isAdmin: boolean): Promise<UserEntity> {
        const newUser = this.userRepository.create({
            email: data.email,
            password: data.password,
            fullName: data.fullName,
            role: isAdmin ? Constants.ADMIN_ROLE : Constants.USER_ROLE
        });
        try {
            const result = await this.userRepository.save(newUser);
            return result;
        } catch (e) {
            return null;
        }
    }

    async findUserByEmail(email: string): Promise<UserEntity> {
        return await this.userRepository.findOne({ email })
    }
}