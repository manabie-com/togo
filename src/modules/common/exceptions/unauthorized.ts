import { UnauthorizedException } from '@nestjs/common';

export class TokenExpiredException extends UnauthorizedException {
  constructor() {
    super('The access token is expired', 'token_expired');
  }
}

export class InvalidTokenException extends UnauthorizedException {
  constructor() {
    super('The access token is invalid', 'invalid_token');
  }
}

export class AuthenticationFailedException extends UnauthorizedException {
  constructor() {
    super('Authentication got failed', 'authentication_failed');
  }
}

export class LoginFailedException extends UnauthorizedException {
  constructor() {
    super('password or username are incorrect', 'login_failed');
  }
}
