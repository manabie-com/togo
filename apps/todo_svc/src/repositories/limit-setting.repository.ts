import {inject} from '@loopback/core';
import {DefaultCrudRepository} from '@loopback/repository';
import {MongodbDataSource} from '../datasources';
import {LimitSetting, LimitSettingRelations} from '../models';

export class LimitSettingRepository extends DefaultCrudRepository<
  LimitSetting,
  typeof LimitSetting.prototype.name,
  LimitSettingRelations
> {
  constructor(@inject('datasources.mongodb') dataSource: MongodbDataSource) {
    super(LimitSetting, dataSource);
  }
}
