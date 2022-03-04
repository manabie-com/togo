import {
  inject,
  injectable,
  Interceptor,
  InvocationContext,
  InvocationResult,
  Provider,
  ValueOrPromise,
} from '@loopback/core';
import {Filter, repository} from '@loopback/repository';
import {LimitSettingRepository, TodoRepository} from '../repositories';
import {startOfDay, endOfDay} from 'date-fns';
// import { User } from '@loopback/authentication-jwt';
import {SecurityBindings, UserProfile} from '@loopback/security';
/**
 * This class will be bound to the application as an `Interceptor` during
 * `boot`
 */
import {LimitSetting, SettingType} from '../models';
const TODO_LIMIT_PER_DAY = 5;
@injectable({tags: {key: ValidateTodoLimitInterceptor.BINDING_KEY}})
export class ValidateTodoLimitInterceptor implements Provider<Interceptor> {
  static readonly BINDING_KEY = `interceptors.${ValidateTodoLimitInterceptor.name}`;

  constructor(
    @repository(TodoRepository)
    public todoRepository: TodoRepository,
    @inject(SecurityBindings.USER)
    private currentUserProfile: UserProfile,
    @repository(LimitSettingRepository)
    private limitSettingRepository: LimitSettingRepository,
  ) {}

  /**
   * This method is used by LoopBack context to produce an interceptor function
   * for the binding.
   *
   * @returns An interceptor function
   */
  value() {
    return this.intercept.bind(this);
  }

  /**
   * The logic to intercept an invocation
   * It checks the area code of the phone number to make sure it matches
   * the provided city name.
   * @param invocationCtx - Invocation context
   * @param next - A function to invoke next interceptor or the target method
   */
  async intercept(
    invocationCtx: InvocationContext,
    next: () => ValueOrPromise<InvocationResult>,
  ) {
    const userId = this.currentUserProfile?.id;
    if (invocationCtx.methodName === 'create' && userId) {
      const today = new Date();
      const todoCount = await this.todoRepository.count({
        userId,
        and: [
          {
            createdAt: {gte: startOfDay(today)},
          },
          {
            createdAt: {lte: endOfDay(today)},
          },
        ],
      });
      const filter = {
        name: SettingType.TODO_LIMIT,
      };
      const todoLimitPerDay: LimitSetting | null =
        (await this.limitSettingRepository.findOne(
          filter as Filter,
        )) as LimitSetting;
      const maxPerDay = todoLimitPerDay?.value ?? TODO_LIMIT_PER_DAY;

      if (todoCount.count >= maxPerDay) {
        const err: ValidationError = new ValidationError(
          `You can not create more task for this user. You are only able to create ${maxPerDay} tasks per day.`,
        );
        err.statusCode = 400;
        throw err;
      }
    }

    const result = await next();
    // Add post-invocation logic here
    return result;
  }
}

class ValidationError extends Error {
  code?: string;
  statusCode?: number;
}
