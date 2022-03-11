import {Entity, model, property} from '@loopback/repository';

@model({settings: {strict: false}})
export class GenericLog extends Entity {
  @property({
    type: 'string',
    id: true,
    generated: false,
    required: true,
  })
  id: string;

  @property({
    type: 'object',
  })
  payload?: object;

  // Define well-known properties here

  // Indexer property to allow additional data
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [prop: string]: any;

  constructor(data?: Partial<GenericLog>) {
    super(data);
  }
}

export interface GenericLogRelations {
  // describe navigational properties here
}

export type GenericLogWithRelations = GenericLog & GenericLogRelations;
