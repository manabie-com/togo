import { NotFoundException } from '@nestjs/common';

export class TaskIsNotFoundException extends NotFoundException {
  constructor(message: string = 'Task is not found.', description: string = 'taskIsNotFound') {
    super([message], description);
  }
}
