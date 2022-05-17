import { IsEmail, IsNotEmpty, IsString } from 'class-validator';

export class AuthSignupDto {
    @IsEmail()
    email: string;

    @IsNotEmpty()
    @IsString()
    password: string;

    @IsNotEmpty()
    @IsString()
    confirmPassword: string;
}
