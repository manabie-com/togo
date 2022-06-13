import db from "../providers/typeorm";
import { TodoRepository } from "../repositories/TodoRepository";

export class TodoRepositoryFactory {
  static async createInstance() {
    const connection = await db.getConnection();
    return connection.getCustomRepository(TodoRepository);
  }
}
