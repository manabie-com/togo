import { Request } from 'express';
import { CanActivate, ExecutionContext, HttpStatus, Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { UserRole } from '../features/user/enum/role.enum';

@Injectable()
export class AdminRoleGuard implements CanActivate {

    constructor(
        private readonly jwtService: JwtService,
    ) {}

    canActivate(context: ExecutionContext): boolean {
        const request: Request = context.switchToHttp().getRequest();

        const authheader = request.header('Authorization');
        const token = authheader && authheader.split(" ")[1];

        try {
            const decodedToken = this.jwtService.decode(token);
            return decodedToken && decodedToken['role'] === UserRole.ADMIN;
        } catch (error) {
            throw new UnauthorizedException({ statusCode: HttpStatus.UNAUTHORIZED ,msg: error.message });
        }
    }
}