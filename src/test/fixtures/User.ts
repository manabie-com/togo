import { DEFAULT_USER_PASSWORD, ADMIN_ROLE_ID, USER_ROLE_ID } from '@modules/common/constant';
import { createPasswordHash } from '@modules/common/utils';
import { User } from '@modules/users/entities/user.entity';

import { faker } from '@faker-js/faker';

const data: User[] = [
  new User({
    id: faker.datatype.uuid(),
    username: 'admin',
    email: 'admin@gmail.com',
    password: createPasswordHash(DEFAULT_USER_PASSWORD),
    displayName: `Admin ${faker.name.findName()}`,
    roleId: ADMIN_ROLE_ID,
  }),
  new User({
    id: faker.datatype.uuid(),
    username: 'user1',
    email: 'user1@gmail.com',
    password: createPasswordHash(DEFAULT_USER_PASSWORD),
    displayName: `User ${faker.name.findName()}`,
    roleId: USER_ROLE_ID,
  }),
  new User({
    id: faker.datatype.uuid(),
    username: 'user2',
    email: 'user2@gmail.com',
    password: createPasswordHash(DEFAULT_USER_PASSWORD),
    displayName: `User ${faker.name.findName()}`,
    roleId: USER_ROLE_ID,
  }),
];

export = data;
