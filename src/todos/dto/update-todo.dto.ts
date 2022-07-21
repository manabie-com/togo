import { OmitType, PartialType } from '@nestjs/swagger';
import { CreateTodoDto } from './create-todo.dto';

export class UpdateTodoDto extends PartialType(OmitType(CreateTodoDto, ['date'])) { }
