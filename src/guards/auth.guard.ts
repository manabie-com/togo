import { Request } from 'express';
import { CanActivate, ExecutionContext, HttpStatus, Injectable, UnauthorizedException } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { AuthService } from '../features/auth/auth.service';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(
        private readonly reflector: Reflector,
        private readonly authService: AuthService,
    ) {}

    canActivate(context: ExecutionContext): boolean {
        const isPublic = this.reflector.get<boolean>( "isPublic", context.getHandler() );
        if ( isPublic ) {
			return true;
		}

        const request: Request = context.switchToHttp().getRequest();

        const authheader = request.header('Authorization');
        const token = authheader && authheader.split(" ")[1];

        try {
            return this.authService.verifyToken(token);
        } catch (error) {
            throw new UnauthorizedException({ statusCode: HttpStatus.UNAUTHORIZED ,msg: error.message });
        }
    }
}