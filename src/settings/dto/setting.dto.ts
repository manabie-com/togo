import { ApiProperty } from "@nestjs/swagger";
import { Type } from "class-transformer";
import { IsNumber, IsString } from "class-validator";
import { UserDto } from "src/users/dto/user.dto";

export class SettingDto {
  @ApiProperty()
  @IsString()
  id: string;

  @ApiProperty({ type: UserDto })
  @Type(() => UserDto)
  user: UserDto;

  @ApiProperty()
  @IsNumber()
  todoPerday: number;
}
