import 'reflect-metadata';
import { Resolver, Query, Mutation, Args } from '@nestjs/graphql';
import { hash, compare } from 'bcrypt';
import { sign } from 'jsonwebtoken';
import { Inject } from '@nestjs/common';
import { User, UserToken } from './entities/user.entity';
import { CreateUserInput } from './dtos/create-user.input';
import { LoginUserInput } from './dtos/login-user.input';
import { PrismaService } from './prisma.service';

@Resolver(() => User)
export class UserResolver {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {}

  @Mutation((returns) => UserToken)
  async signup(
    @Args('createUserInput') createUserInput: CreateUserInput,
  ): Promise<UserToken> {
    const hashedPassword = await hash(createUserInput.password, 10);

    const created = await this.prismaService.user.create({
      data: {
        email: createUserInput.email,
        password: hashedPassword,
        name: createUserInput.name,
        maxJob: createUserInput.maxJob,
      },
    });

    return {
      ...created,
      token: sign({ userId: created.id }, process.env.APP_SECRET),
    };
  }

  @Mutation((returns) => UserToken)
  async login(@Args('loginUserInput') loginUserInput: LoginUserInput) {
    const user = await this.prismaService.user.findUnique({
      where: {
        email: loginUserInput.email,
      },
    });

    if (!user) {
      throw new Error(`No user found with email: ${loginUserInput.email}`);
    }

    const passwordParse = await compare(loginUserInput.password, user.password);

    if (!passwordParse) {
      throw new Error('Invalid password');
    }

    return {
      ...user,
      token: sign({ userId: user.id }, process.env.APP_SECRET),
    };
  }
}
