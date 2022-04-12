import { DynamicModule, Module } from '@nestjs/common';
import { TypeOrmModule, TypeOrmModuleOptions } from '@nestjs/typeorm';
import { PgSubModule } from './pg-sub.module';
import { PgService } from './services/pg.service';

@Module({})
export class PgModule {
  static register(config: {
    entities: any[];
    synchronize?: boolean;
  }): DynamicModule {
    if (typeof config.synchronize !== 'boolean') {
      config.synchronize = true;
    }
    return {
      module: PgModule,
      imports: [
        TypeOrmModule.forRootAsync({
          imports: [PgSubModule],
          useFactory: async (pgService: PgService) =>
            ({
              ...pgService.getTypeOrmConfig({ entities: config.entities }),
              synchronize: config.synchronize,
            } as TypeOrmModuleOptions),
          inject: [PgService],
        }),
      ],
    };
  }
}
