import { EntityManager, EntityRepository, Between } from "typeorm";
import { ITodoRepository } from "../core/repository/ITodoRepository";
import { Todo } from "../entity/Todo";
import _ from "lodash";
import moment from "moment";

@EntityRepository()
export class TodoRepository implements ITodoRepository {
  constructor(private manager: EntityManager) {}

  async saveTodo(params) {
    if (_.isEmpty(params)) {
      return {};
    }

    const todo = new Todo();
    const data = _.assign(todo, params);

    return await this.manager.save(data);
  }

  async getCurrentTasksByUserId(userId) {
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
