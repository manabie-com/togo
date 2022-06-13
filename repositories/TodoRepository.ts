import { EntityManager, EntityRepository, Equal } from "typeorm";
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
    const filters = {
      userId,
      creationDate: moment().format("YYYY-MM-DD"),
    };
    const [data] = await this.manager
      .getRepository(Todo)
      .createQueryBuilder()
      .select("COUNT(id)", "tasks")
      .andWhere("user_id = :userId")
      .andWhere("DATE(creation_date) = :creationDate")
      .setParameters(filters)
      .execute();

    const { tasks } = data;

    return tasks;
  }
}
