import 'reflect-metadata';
import { ObjectType, Field, ID, Int } from '@nestjs/graphql';
import { User } from './user.entity';

@ObjectType()
export class Task {
  @Field((type) => Int)
  id: number;

  @Field((type) => Date)
  createdAt: Date;

  @Field((type) => Date)
  updatedAt: Date;

  @Field((type) => String)
  title: string;

  @Field((type) => String)
  content: string;

  @Field((type) => User, { nullable: true })
  user?: User | null;

  @Field((type) => Int, { nullable: true })
  userId: number | null;
}
