import { IsString, IsNotEmpty, IsEnum } from 'class-validator';
import { TodoStatus } from '../todo/todo.entity';

export class UpdateTodoStatusDto {
  @IsNotEmpty()
  @IsString()
  id: string;

  @IsNotEmpty()
  @IsEnum(TodoStatus)
  status: TodoStatus;
}
