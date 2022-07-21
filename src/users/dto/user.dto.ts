import { ApiProperty } from "@nestjs/swagger";
import { IsString, Matches, MaxLength } from "class-validator";

export class UserDto {
  @ApiProperty()
  id: string;

  @ApiProperty({
    pattern: "^[a-z][a-z0-9_]{4,9}$"
  })
  @Matches(/^[a-z][a-z0-9_]{4,9}$/)
  username: string;

  @ApiProperty()
  @IsString()
  @MaxLength(255)
  firstName: string;

  @ApiProperty()
  @IsString()
  @MaxLength(255)
  lastName: string;
}
