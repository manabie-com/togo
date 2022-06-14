import { ITodo, ITodoParams } from "../models/ITodo";

export interface ITodoRepository {
  /**
   *
   * @param params
   */
  saveTodo(params: ITodoParams): Promise<ITodo>;

  /**
   *
   * @param userId
   */
  getCurrentTasksByUserId(userId: number): Promise<number>;
}
