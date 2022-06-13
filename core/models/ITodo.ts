export interface ITodo {
  id: number;
  task: string;
  userId: number;
  creationDate?: string;
}

export interface ITodoParams {
  task: string;
  userId: number;
}
