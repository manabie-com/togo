import { EntityManager, EntityRepository } from "typeorm";
import { ITodoRepository } from "../core/repository/ITodoRepository";
import { Todo } from "../entity/Todo";
import _ from "lodash";

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
}
