import { IsNotEmpty, IsNumber } from 'class-validator'

export class TakeTaskDTO {
  @IsNotEmpty()
  @IsNumber()
  taskId: number
}
