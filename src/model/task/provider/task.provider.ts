import { TASK } from "src/constance/variable";
import { Task } from "../schema/task.entity";

export const taskProviders = [
  {
    provide: TASK,
    useValue: Task,
  },
];