import { Injectable } from '@nestjs/common';
import { TypeOrmModuleOptions } from '@nestjs/typeorm';
import { EnvironmentService } from '../../environment/services/environment.service';
import { PgConstants } from '../constants/pg.constant';

@Injectable()
export class PgService {
  constructor(private readonly environmentService: EnvironmentService) {}

  public getTypeOrmConfig(config?: { entities: any[] }): TypeOrmModuleOptions {
    const connection = {
      host: this.environmentService.getKey(PgConstants.POSTGRES_HOST),
      port: parseInt(this.environmentService.getKey(PgConstants.POSTGRES_PORT)),
      username: this.environmentService.getKey(PgConstants.POSTGRES_USER),
      password: this.environmentService.getKey(PgConstants.POSTGRES_PASSWORD),
      database: this.environmentService.getKey(PgConstants.POSTGRES_DATABASE),
    };

    if (process.env.NODE_ENV === 'test') {
      return {
        type: 'sqlite',
        database: 'test.db',
        dropSchema: true,
        entities: config.entities,
        synchronize: true,
        logging: false,
      };
    }

    return {
      type: PgConstants.POSTGRES,
      ...connection,
      entities: config?.entities || [PgConstants.TYPE_ORM_ENTITIES],
      migrationsTableName: PgConstants.TYPE_ORM_MIGRATION_TABLE_NAME,
      migrations: [PgConstants.TYPE_ORM_MIGRATIONS],
      extra: {
        max: 100,
      },
      synchronize: true,
      logging: true,
    } as TypeOrmModuleOptions;
  }
}
