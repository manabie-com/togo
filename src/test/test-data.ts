import { ADMIN_ROLE_ID, ADMIN_ROLE_NAME, USER_ROLE_ID, USER_ROLE_NAME } from '@modules/common/constant';
import { PermissionAction } from '@modules/permissions/permission.action.enum';
import { PermissionResource } from '@modules/permissions/permission.resource.enum';

import { faker } from '@faker-js/faker';

const roles = [
  {
    id: ADMIN_ROLE_ID,
    name: ADMIN_ROLE_NAME,
  },
  {
    id: USER_ROLE_ID,
    name: USER_ROLE_NAME,
  },
];

const permissions = [
  {
    id: faker.datatype.uuid(),
    name: 'Create User',
    resource: PermissionResource.Users,
    action: PermissionAction.Create,
  },
  {
    id: faker.datatype.uuid(),
    name: 'Update User',
    resource: PermissionResource.Users,
    action: PermissionAction.Update,
  },
  { id: faker.datatype.uuid(), name: 'List User', resource: PermissionResource.Users, action: PermissionAction.List },
  {
    id: faker.datatype.uuid(),
    name: 'Detail User',
    resource: PermissionResource.Users,
    action: PermissionAction.Detail,
  },
  {
    id: faker.datatype.uuid(),
    name: 'Delete User',
    resource: PermissionResource.Users,
    action: PermissionAction.Delete,
  },
  {
    id: faker.datatype.uuid(),
    name: 'Create Task',
    resource: PermissionResource.Tasks,
    action: PermissionAction.Create,
  },
  {
    id: faker.datatype.uuid(),
    name: 'Update Task',
    resource: PermissionResource.Tasks,
    action: PermissionAction.Update,
  },
  { id: faker.datatype.uuid(), name: 'List Task', resource: PermissionResource.Tasks, action: PermissionAction.List },
  {
    id: faker.datatype.uuid(),
    name: 'Detail Task',
    resource: PermissionResource.Tasks,
    action: PermissionAction.Detail,
  },
  {
    id: faker.datatype.uuid(),
    name: 'Delete Task',
    resource: PermissionResource.Tasks,
    action: PermissionAction.Delete,
  },
  { id: faker.datatype.uuid(), name: 'Pick Task', resource: PermissionResource.Tasks, action: PermissionAction.Pick },
];

const baseActions = [
  PermissionAction.Create,
  PermissionAction.Update,
  PermissionAction.List,
  PermissionAction.Detail,
  PermissionAction.Delete,
];

const taskActions = [PermissionAction.Pick, PermissionAction.List];

const fullActions = [...baseActions];

const mapRolePermission = [
  {
    roleName: ADMIN_ROLE_NAME,
    permissions: [
      { resource: PermissionResource.Users, actions: fullActions },
      { resource: PermissionResource.Tasks, actions: fullActions },
    ],
  },
  {
    roleName: USER_ROLE_NAME,
    permissions: [{ resource: PermissionResource.Tasks, actions: taskActions }],
  },
];

const rolePermissions = [];

mapRolePermission.forEach((x) => {
  const role = roles.find((role) => role.name == x.roleName);

  if (role) {
    x.permissions.forEach((y) => {
      y.actions.forEach((act) => {
        const permission = permissions.find((per) => per.resource == y.resource && act === per.action);

        if (permission) {
          const rolePermission = { roleId: role.id, permissionId: permission.id };

          rolePermissions.push(rolePermission);
        }
      });
    });
  }
});

export { roles, permissions, rolePermissions };
