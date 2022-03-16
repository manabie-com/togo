import 'reflect-metadata';
import {
  Resolver,
  Query,
  Mutation,
  Args,
  Context,
  ResolveField,
  Root,
  InputType,
  Field,
  registerEnumType,
} from '@nestjs/graphql';
import { Inject } from '@nestjs/common';
import { User } from './entities/user.entity';
import { Task } from './entities/task.entity';
import { CreateTaskInput } from './dtos/create-task.input';
import { PrismaService } from './prisma.service';
import { getUserId } from './context';

@InputType()
class TaskOrderByUpdatedAtInput {
  @Field((type) => SortOrder)
  updatedAt: SortOrder;
}

enum SortOrder {
  asc = 'asc',
  desc = 'desc',
}

registerEnumType(SortOrder, {
  name: 'SortOrder',
});

@Resolver(() => Task)
export class TaskResolver {
  constructor(
    @Inject(PrismaService) private readonly prismaService: PrismaService,
  ) {}

  @ResolveField()
  user(@Root() task: Task): Promise<User | null> {
    return this.prismaService.task
      .findUnique({
        where: {
          id: task.id,
        },
      })
      .user();
  }

  @Query((returns) => Task, { nullable: true })
  async taskById(@Args('id') id: number): Promise<Task> {
    return await this.prismaService.task.findUnique({
      where: { id },
    });
  }

  @Query((returns) => [Task])
  async tasks(
    @Args('searchString', { nullable: true }) searchString: string,
    @Args('skip', { nullable: true }) skip: number,
    @Args('take', { nullable: true }) take: number,
    @Args('orderBy', { nullable: true }) orderBy: TaskOrderByUpdatedAtInput,
    @Context() ctx,
  ): Promise<Task[]> {
    const or = searchString
      ? {
          OR: [
            { title: { contains: searchString } },
            { content: { contains: searchString } },
          ],
        }
      : {};

    return await this.prismaService.task.findMany({
      where: {
        ...or,
      },
      take: take || undefined,
      skip: skip || undefined,
      orderBy: orderBy || undefined,
    });
  }

  @Mutation((returns) => Task)
  async createTask(
    @Args('createTaskInput') createTaskInput: CreateTaskInput,
    @Context() ctx,
  ): Promise<Task> {
    const userId = getUserId(ctx);

    if (!userId) {
      throw new Error('Could not authenticate user.');
    }

    const findMaxJob = await this.prismaService.user.findUnique({
      where: { id: userId },
      select: {
        maxJob: true,
      },
    });

    if (!findMaxJob?.maxJob) {
      throw new Error('Could not get max job.');
    }

    const listTasks =
      (await this.prismaService.task.findMany({
        where: { userId },
      })) ?? [];

    if (listTasks?.length >= findMaxJob?.maxJob) {
      throw new Error('User reach task limit error.');
    }

    return this.prismaService.task.create({
      data: {
        title: createTaskInput.title,
        content: createTaskInput.content,
        user: {
          connect: {
            id: userId,
          },
        },
      },
      include: {
        user: true,
      },
    });
  }
}
