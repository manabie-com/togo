import { IUser } from "../models/IUser";

export interface IUserRepository {
  getUserById(): Promise<IUser | boolean>;
}
