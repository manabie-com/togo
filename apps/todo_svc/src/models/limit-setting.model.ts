import {User, UserWithRelations} from '@loopback/authentication-jwt';
import {belongsTo, Entity, model, property} from '@loopback/repository';

export enum SettingType {
  TODO_LIMIT = 'TODO_LIMIT',
}

@model({settings: {strict: false}})
export class LimitSetting extends Entity {
  @property({
    type: 'string',
    id: true,
    generated: false,
    mongodb: {dataType: 'ObjectID'},
  })
  id?: string;

  @property({
    type: 'string',
    required: true,
    id: 1,
    jsonSchema: {
      enum: Object.values(SettingType),
    },
    default: SettingType.TODO_LIMIT,
  })
  name: string;

  @property({
    type: 'number',
    required: true,
  })
  value: number;

  @belongsTo(() => User)
  // @property({
  //   type: 'string',
  //   id: 2
  // })
  userId: string;

  // Define well-known properties here

  // Indexer property to allow additional data
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [prop: string]: any;

  constructor(data?: Partial<LimitSetting>) {
    super(data);
  }
}

export interface LimitSettingRelations {
  // describe navigational properties here
  user?: UserWithRelations;
}

export type LimitSettingWithRelations = LimitSetting & LimitSettingRelations;
