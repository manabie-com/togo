import { PartialType } from '@nestjs/swagger';
import { CreateTodoTaskDto } from './create-todo-task.dto';

export class UpdateTodoTaskDto extends PartialType(CreateTodoTaskDto) {}
