import db from "../providers/typeorm";
import { UserRepository } from "../repositories/UserRepository";

export class UserRepositoryFactory {
  static async createInstance() {
    const connection = await db.getConnection();
    return connection.getCustomRepository(UserRepository);
  }
}
