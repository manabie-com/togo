import { IsEmail, IsNotEmpty } from 'class-validator';

export class AuthSigninDto {
    @IsEmail()
    email: string;

    @IsNotEmpty()
    password: string;
}