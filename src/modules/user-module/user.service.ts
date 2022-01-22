import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { compareSync } from 'bcrypt';
import { Checker } from 'src/interfaces/checker.interface';
import { Constants } from '../../utils/constants';
import { ConfigureMaxTaskDTO } from './dto/configure_max_task.dto';
import { FindUserDTO } from './dto/find.dto';
import { LoginDTO } from './dto/login.dto';
import { RegisterDTO } from './dto/register.dto';
import { TaskDTO } from './dto/task.dto';
import { UserRepository } from './user.repository';

@Injectable()
export class UserService {
    constructor(
        private readonly userRepository: UserRepository,
        private readonly jwtService: JwtService,
    ) { }

    async register(data: RegisterDTO, isAdmin: boolean = false): Promise<Checker> {
        const result = await this.userRepository.register(data, isAdmin);
        if (result) {
            const data = {
                id: result.id,
                email: result.email,
                role: result.role
            };
            return {
                isSuccess: true,
                data: {
                    token: this.jwtService.sign(data, {
                        expiresIn: '30d',
                    }),
                    email: result.email,
                },
            };
        }
        return Constants.FAIL_CHECK
    }

    verifyAccount(token: string): Checker {
        try {
            const validToken = this.jwtService.verify(token);
            return {
                isSuccess: true,
                data: validToken,
            };
        } catch (e) {
            return {
                isSuccess: false,
                data: null,
            };
        }

    }

    async login(data: LoginDTO): Promise<Checker> {
        const user = await this.userRepository.findUserByEmail(data.email);
        if (user) {
            const hashPassword = user.password;
            const isCorrectPassword = compareSync(data.password, hashPassword);
            if (isCorrectPassword) {
                const data = {
                    id: user.id,
                    email: user.email,
                    role: user.role
                };
                return {
                    isSuccess: true,
                    data: {
                        token: this.jwtService.sign(data, {
                            expiresIn: '30d',
                        }),
                        email: user.email,
                    },
                };
            }
        }
        return Constants.FAIL_CHECK
    }

    async find(query: FindUserDTO): Promise<Checker> {
        const defaultValue = {
            pageIndex: 1,
            perPage: 10
        }
        const result = await this.userRepository.find({
            ...defaultValue,
            ...query
        });
        if (result) {
            return {
                isSuccess: true,
                data: result
            };
        }
        return Constants.FAIL_CHECK
    }

    async addTask(userId: number, task: TaskDTO): Promise<Checker> {
        let res = await this.userRepository.addNewTask(userId, task);
        if (res) {
            return Constants.SUCCESS_CHECK;
        }
        return Constants.FAIL_CHECK;
    }

    async configureMaxTask(config: ConfigureMaxTaskDTO): Promise<Checker>{
        let res = await this.userRepository.configureMaxTask(config);
        if (res) {
            return Constants.SUCCESS_CHECK;
        }
        return Constants.FAIL_CHECK;
    }
}
