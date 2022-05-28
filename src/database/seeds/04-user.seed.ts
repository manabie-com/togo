import { Factory, Seeder } from 'typeorm-seeding';
import { faker } from '@faker-js/faker';
import { Connection } from 'typeorm';
import { createPasswordHash } from '@modules/common/utils';
import { USER_ROLE_ID, ADMIN_ROLE_ID, DEFAULT_USER_PASSWORD } from '@modules/common/constant';
import { cleanTable } from '@modules/common/utils/query';

const users = [
  {
    username: 'admin',
    email: 'admin@gmail.com',
    password: createPasswordHash(DEFAULT_USER_PASSWORD),
    displayName: `Admin ${faker.name.findName()}`,
    roleId: ADMIN_ROLE_ID,
  },
  {
    username: 'user1',
    email: 'user1@gmail.com',
    password: createPasswordHash(DEFAULT_USER_PASSWORD),
    displayName: `User ${faker.name.findName()}`,
    roleId: USER_ROLE_ID,
  },
  {
    username: 'user2',
    email: 'user2@gmail.com',
    password: createPasswordHash(DEFAULT_USER_PASSWORD),
    displayName: `User ${faker.name.findName()}`,
    roleId: USER_ROLE_ID,
  },
];

export default class CreateUserSeeder implements Seeder {
  tableName = 'users';

  public async run(factory: Factory, connection: Connection): Promise<any> {
    await cleanTable(connection, this.tableName);

    await connection.createQueryBuilder().insert().into(this.tableName).values(users).execute();
  }
}
