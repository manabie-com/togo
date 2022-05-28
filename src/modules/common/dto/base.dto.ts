import { Allow } from 'class-validator';

export class BaseDto {
  @Allow()
  context?: {
    params?: any;
    query?: any;
    headers?: any;
  };
}
