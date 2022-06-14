import { IUser } from "../models/IUser";

export interface IUserRepository {
  /**
   *
   * @param userId
   */
  getUserById(userId): Promise<IUser | boolean>;
}
