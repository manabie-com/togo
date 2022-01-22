import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { compareSync } from 'bcrypt';
import { Checker } from 'src/interfaces/checker.interface';
import { Constants } from '../../utils/constants';
import { LoginDTO } from './dto/login.dto';
import { RegisterDTO } from './dto/register.dto';
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

}
