import { ITodo, ITodoParams } from "../models/ITodo";

export interface ITodoRepository {
  saveTodo(params: ITodoParams, headers?: any): Promise<ITodo>;
}
