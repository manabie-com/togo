import {inject} from '@loopback/core';
import {DefaultCrudRepository} from '@loopback/repository';
import {MongodbDataSource} from '../datasources';
import {AccessLog, AccessLogRelations} from '../models';

export class AccessLogRepository extends DefaultCrudRepository<
  AccessLog,
  typeof AccessLog.prototype.id,
  AccessLogRelations
> {
  constructor(@inject('datasources.mongodb') dataSource: MongodbDataSource) {
    super(AccessLog, dataSource);
  }
}
