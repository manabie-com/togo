import { JwtOptionsFactory, JwtModuleOptions } from '@nestjs/jwt';
import { Injectable } from '@nestjs/common';

@Injectable()
export class JwtConfig implements JwtOptionsFactory {
  createJwtOptions(): JwtModuleOptions {
    return {
      signOptions: {
        expiresIn: process.env.JWT_EXPIRE_IN,
        algorithm: 'RS256',
      },
      privateKey: process.env.JWT_PRIVATE_KEY,
      publicKey: process.env.JWT_PUBLIC_KEY,
      verifyOptions: { algorithms: ['RS256'] },
    };
  }
}
