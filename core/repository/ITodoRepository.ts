import { ITodo, ITodoParams } from "../models/ITodo";

export interface ITodoRepository {
  saveTodo(params: ITodoParams): Promise<ITodo>;
  getCurrentTasksByUserId(userId: number): Promise<number>;
}
