import {Entity, model, property} from '@loopback/repository';

@model()
export class AccessLog extends Entity {
  @property({
    type: 'date',
    required: true,
  })
  createdAt: string;

  @property({
    type: 'string',
    required: true,
  })
  serviceName: string;

  @property({
    type: 'string',
    id: true,
    generated: true,
  })
  id?: string;

  @property({
    type: 'string',
    required: true,
  })
  url: string;

  @property({
    type: 'string',
    required: true,
  })
  method: string;

  @property({
    type: 'object',
  })
  payload?: object;

  constructor(data?: Partial<AccessLog>) {
    super(data);
  }
}

export interface AccessLogRelations {
  // describe navigational properties here
}

export type AccessLogWithRelations = AccessLog & AccessLogRelations;
