import { Role } from '@modules/roles/entities/role.entity';
import { roles } from '@test/test-data';

const data: Role[] = roles.map(
  (role) =>
    new Role({
      id: role.id,
      name: role.name,
    }),
);

export = data;
