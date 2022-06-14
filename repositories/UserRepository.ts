import { EntityManager, EntityRepository } from "typeorm";
import { User } from "../entity/User";

@EntityRepository()
export class UserRepository {
  constructor(private manager: EntityManager) {}

  /**
   *
   * @param userId
   * @returns user
   */
  async getUserById(userId: number) {
    return this.manager.findOne(User, {
      where: { userId },
    });
  }
}
