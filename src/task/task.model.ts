import { Schema, model } from 'mongoose';
import { TaskStatusEnum } from './task.enum';

const reasonSchema = new Schema({
  errorCode: {
    type: String
  },
  message: {
    type: String
  }
});

const taskSchema = new Schema({
  name: { type: String, required: true },
  userId: { type: String, required: true },
  status: {
    type: String,
    required: true,
    enum: [...Object.values(TaskStatusEnum)],
    default: TaskStatusEnum.PENDING
  },
  reason: { type: reasonSchema }
});

const taskModel = model('task', taskSchema);

export default taskModel;
