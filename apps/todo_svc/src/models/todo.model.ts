import { User, UserWithRelations } from '@loopback/authentication-jwt';
import {belongsTo, Entity, model, property} from '@loopback/repository';

@model()
export class Todo extends Entity {
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
    jsonSchema: {
      maxLength: 200,
      errorMessage: 'name must be less than 200 characters',
    },
  })
  title: string;

  @property({
    type: 'string',
    jsonSchema: {
      maxLength: 500,
      errorMessage: 'name must be less than 500 characters',
    },
  })
  desc?: string;

  @property({
    type: 'boolean',
  })
  isComplete?: boolean;

  @property({
    type: 'date',
    defaultFn: "now",
  })
  createdAt?: Date;


  @belongsTo(() => User)
  userId: string;


  constructor(data?: Partial<Todo>) {
    super(data);
  }
}

export interface TodoRelations {
  // describe navigational properties here
  user?: UserWithRelations;
}

export type TodoWithRelations = Todo & TodoRelations;
