import { LIMIT_TASK } from "src/constance/variable";
import { LimitTask } from "../schema/limitTask.entity";

export const limitTaskProviders = [
  {
    provide: LIMIT_TASK,
    useValue: LimitTask,
  },
];