import { EntityManager, EntityRepository, Between } from "typeorm";
import { ITodoRepository } from "../core/repository/ITodoRepository";
import { Todo } from "../entity/Todo";
import _ from "lodash";
import moment from "moment";
import { ITodoParams } from "../core/models/ITodo";

@EntityRepository()
export class TodoRepository implements ITodoRepository {
  constructor(private manager: EntityManager) {}

  /**
   *
   * @param params
   * @returns created record
   */
  async saveTodo(params: ITodoParams) {
    if (_.isEmpty(params)) {
      return {};
    }

    const todo = new Todo();
    const data = _.assign(todo, params);

    return await this.manager.save(data);
  }

  /**
   *
   * @param userId
   * @returns a count of user's current day tasks
   */
  async getCurrentTasksByUserId(userId: number) {
    const startOfDay = moment().startOf("day").toISOString();
    const endOfDay = moment().endOf("day").toISOString();

    const [, count] = await this.manager.findAndCount(Todo, {
      where: {
        userId,
        creationDate: Between(startOfDay, endOfDay),
      },
    });

    return count;
  }
}
