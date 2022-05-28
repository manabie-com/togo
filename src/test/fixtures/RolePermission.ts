import { rolePermissions } from '@test/test-data';

const data: Record<string, any>[] = rolePermissions.map((item) => {
  return {
    roleId: item.roleId,
    permissionId: item.permissionId,
  };
});
export = data;
