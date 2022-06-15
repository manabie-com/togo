import { TASK } from '../../../constance/variable';
import { Task } from '../schema/task.entity';

export const taskProviders = [
  {
    provide: TASK,
    useValue: Task,
  },
];
