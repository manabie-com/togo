import { Factory, Seeder } from 'typeorm-seeding'
import { Connection } from 'typeorm'
import { UserEntity } from '../../features/user/entities/user.entity'
import { CryptoUtil } from '../../utils/crypto.util';
import { UserRole } from 'src/features/user/enum/role.enum';
 
export default class CreateUsers implements Seeder {
  public async run(factory: Factory, connection: Connection): Promise<any> {
    await connection
      .createQueryBuilder()
      .insert()
      .into(UserEntity)
      .values([
        { 
            email: 'admin@gmail.com', 
            password: await CryptoUtil.hash('admin123'),
            role: UserRole.ADMIN
        },
      ])
      .execute()
  }
}