import { Injectable, NotFoundException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { CryptoUtil } from '../../utils/crypto.util';
import { UserService } from '../user/user.service';
import { AuthSigninDto } from './dto/auth-signin.dto';

@Injectable()
export class AuthService {
    constructor(
		private readonly userService: UserService,
		private readonly jwtService: JwtService
	){}

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

		return await this.jwtService.sign({ email, role: user.role });
	}

	verifyToken(token: string): boolean {
		const isTokenValid = this.jwtService.verify(token);
		if(!isTokenValid) {
			return false;
		}

		return true;
	}
}
