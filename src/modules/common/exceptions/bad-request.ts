import { BadRequestException } from '@nestjs/common';

export class TokenInvalidOrExpiredBadRequestException extends BadRequestException {
  constructor(message: string = 'tokenIsInvalidOrExpired.', description: string = 'invalidToken') {
    super([message], description);
  }
}

export class TaskIsAssignedToYouBadRequestException extends BadRequestException {
  constructor(message: string = 'This task is assigned to you already.', description: string = 'taskIsAssigned') {
    super([message], description);
  }
}

export class TaskIsHandedByOtherBadRequestException extends BadRequestException {
  constructor(
    message: string = 'Task is handled by the other user.',
    description: string = 'taskIsHandledByOtherUser',
  ) {
    super([message], description);
  }
}

export class ReachedMaximumTaskTodayBadRequestException extends BadRequestException {
  constructor(
    message: string = `You reached a maximum limit of tasks per user that can be added per day.`,
    description: string = 'userReachMaximumTaskToday',
  ) {
    super([message], description);
  }
}
