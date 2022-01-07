import { CreateTaskDto } from "../../src/task/dto/create-task-dto";
import { Task } from "../../src/task/schemas/task.schema";

let createTaskDto = new CreateTaskDto();
createTaskDto.title = 'anytitle';
createTaskDto.content = 'anycontent';
createTaskDto.dateTime = new Date();

export const mockCreateTaskDto = (): CreateTaskDto => ({
  ...createTaskDto
});

export const mockTask = (): Task => ({
  id: 'anyid',
  title: 'anytitle',
  content: 'anycontent',
  dateTime: new Date(),
  createdAt: new Date(),
  updatedAt: new Date()
})