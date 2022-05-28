import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { Reflector } from '@nestjs/core';

import { UserService } from '@modules/users/user.service';

@Injectable()
export class RolesGuard implements CanActivate {
  constructor(private reflector: Reflector, private readonly userService: UserService) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const roles = this.reflector.get<string[]>('roles', context.getHandler());

    if (!roles || !roles.length) {
      return false;
    }

    const request = context.switchToHttp().getRequest();

    return await this.canAccessResource(roles, request?.user);
  }

  async canAccessResource(roles: string[], user: any): Promise<boolean> {
    let canAccess = false;

    await Promise.all(
      roles.map(async (r) => {
        const role: Record<string, any> = (r as unknown) as Record<string, any>;

        const hasPermission = await this.userService.checkPermission(user?.userId, role?.resource, role?.action);

        if (hasPermission) {
          canAccess = true;
        }
      }),
    );

    return canAccess;
  }
}
