import 'reflect-metadata';
import { ObjectType, Field, Int, PartialType } from '@nestjs/graphql';
import { IsEmail } from 'class-validator';

@ObjectType()
export class User {
  @Field((type) => Int)
  id: number;

  @Field()
  @IsEmail()
  email: string;

  @Field((type) => String, { nullable: true })
  name?: string | null;

  @Field((type) => Int)
  maxJob: number;
}

@ObjectType()
export class UserToken extends PartialType(User) {
  @Field(() => String, { nullable: true })
  token: string | null;
}
