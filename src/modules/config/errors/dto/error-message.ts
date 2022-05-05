import { HttpStatus } from '@nestjs/common';
import { IErrorMessages } from './error-message.interface';

export const errorMessagesConfig: { [messageCode: string]: IErrorMessages } = {
  'user:create:missingInformation': {
    type: 'BadRequest',
    httpStatus: HttpStatus.BAD_REQUEST,
    errorMessage: 'Unable to create a new user with missing information.',
    userMessage: 'Unable to create a new user with missing information.',
  },
};
