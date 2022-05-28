import { BaseDto } from '@modules/common/dto/base.dto';
import { ApiProperty, PartialType } from '@nestjs/swagger';
import { ValidateIf, IsNotEmpty, IsEnum, IsOptional } from 'class-validator';

export enum GrantType {
  Password = 'password',
  RefreshToken = 'refresh_token',
}

export class CreateTokenDto extends PartialType(BaseDto) {
  @ApiProperty({ enum: GrantType })
  @IsEnum(GrantType, { each: true })
  readonly grant_type: GrantType;

  @ApiProperty()
  @ValidateIf((o) => o.grantType === 'password')
  @IsNotEmpty()
  readonly username?: string;

  @ApiProperty()
  @ValidateIf((o) => o.grantType === 'password')
  @IsNotEmpty()
  readonly password?: string;

  @ApiProperty()
  @ValidateIf((o) => o.grantType === 'refresh_token')
  @IsNotEmpty()
  readonly refresh_token?: string;

  @ApiProperty()
  @IsOptional()
  readonly rememberMe?: boolean;
}
