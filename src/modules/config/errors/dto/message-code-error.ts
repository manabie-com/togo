import { errorMessagesConfig } from './error-message';
import { IErrorMessages } from './error-message.interface';

export class MessageCodeError extends Error {
  public messageCode: string;
  public httpStatus: number;
  public errorMessage: string;

  constructor(messageCode: string) {
    super();

    const errorMessageConfig = this.getMessageFromMessageCode(messageCode);
    if (!errorMessageConfig)
      throw new Error('Unable to find message code error.');

    Error.captureStackTrace(this, this.constructor);
    this.name = this.constructor.name;
    this.httpStatus = errorMessageConfig.httpStatus;
    this.messageCode = messageCode;
    this.errorMessage = errorMessageConfig.errorMessage;
    this.message = errorMessageConfig.userMessage;
  }

  /**
   * @description: Find the error config by the given message code.
   * @param {string} messageCode
   * @return {IErrorMessages}
   */
  private getMessageFromMessageCode(messageCode: string): IErrorMessages {
    const errorMessageConfig: IErrorMessages | undefined =
      errorMessagesConfig[messageCode] || undefined;

    if (!errorMessageConfig)
      throw new Error('Unable to find the given message code error.');
    return errorMessageConfig;
  }
}
