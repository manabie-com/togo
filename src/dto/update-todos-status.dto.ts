import { IsString, IsNotEmpty, IsEnum } from 'class-validator';
import { TodoStatus } from '../todo/todo.entity';

export class UpdateTodosStatusDto {
  @IsNotEmpty()
  @IsString({ each: true })
  ids: string[];

  @IsNotEmpty()
  @IsEnum(TodoStatus)
  status: TodoStatus;
}
