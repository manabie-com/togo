import { Request } from 'express';
import { CanActivate, ExecutionContext, Injectable, UnauthorizedException } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { AuthService } from 'src/features/auth/auth.service';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(
        private readonly reflector: Reflector,
        private readonly authService?: AuthService,
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
            this.authService.verifyToken(token);
            return true;
        } catch (error) {
            throw new UnauthorizedException();
        }
    }
}