import { Permission } from '@modules/permissions/permission.entity';
import { permissions } from '@test/test-data';

const data: Permission[] = permissions.map((per) => {
  return new Permission({
    id: per.id,
    name: per.name,
    resource: per.resource,
    action: per.action,
  });
});

export = data;
