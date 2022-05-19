import { Global, Module } from "@nestjs/common";
import { ConfigModule, ConfigService } from "@nestjs/config";
import { JwtModule } from "@nestjs/jwt";

@Global()
@Module({
  imports: [
    JwtModule.registerAsync({
        imports: [ConfigModule],
        inject: [ConfigService],
        useFactory: (configService: ConfigService) => ({
            secret: configService.get('AUTH_JWT_SECRET'),
            signOptions: { expiresIn: configService.get('AUTH_JWT_EXPIRED_TIME') }
        }),
    })
  ],
  exports: [JwtModule]
})
export class CoreModule {}