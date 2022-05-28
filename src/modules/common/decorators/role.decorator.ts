import { SetMetadata } from '@nestjs/common';

export const UseRoles = (...roles: Record<string, any>[]) => SetMetadata('roles', roles);
