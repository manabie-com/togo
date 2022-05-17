import { ForbiddenException, Injectable, InternalServerErrorException, NotFoundException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { CryptoUtil } from 'src/utils/crypto.util';
import { UserEntity } from '../user/entities/user.entity';
import { UserService } from '../user/user.service';
import { AuthSigninDto } from './dto/auth-signin.dto';
import { AuthSignupDto } from './dto/auth-signup.dto';

@Injectable()
export class AuthService {
    constructor(
		private readonly userService: UserService,
		private readonly jwtService: JwtService
	){}

	async signup(registerUserDto: AuthSignupDto): Promise<UserEntity>{
		const { email, password } = registerUserDto;

		const isUserExist = await this.userService.findOneByEmail(email);
		if(isUserExist) {
			throw new NotFoundException(`User with email ${email} already existed`);
		}

		if(registerUserDto.password !== registerUserDto.confirmPassword) {
			throw new ForbiddenException(`password and confirmation password is not matched`)
		}

		const hashPassword = await CryptoUtil.hash(password);

		const user: UserEntity = await this.userService.create({ email, password: hashPassword });
		return user;
	}

	async signin(signinUserDto: AuthSigninDto): Promise<string>{
		const { email, password } = signinUserDto;

		const user = await this.userService.findOneByEmail(email);
		if(!user) {
			throw new NotFoundException(`Email or password is invalid. Please try again.`);
		}

		const isValidPassword = await CryptoUtil.compareHashWithPlainText(user.password, password);
		if(!isValidPassword) {
			throw new NotFoundException(`Email or password is invalid. Please try again.`);
		}

		return await this.jwtService.sign({ email });
	}

	verifyToken(token: string): boolean {
		const isTokenValid = this.jwtService.verify(token);
		if(!isTokenValid) {
			return false;
		}

		return true;
	}
}
