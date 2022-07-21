import { ApiProperty } from "@nestjs/swagger";
import { Type } from "class-transformer";
import { IsDateString, IsString } from "class-validator";
import { UserDto } from "src/users/dto/user.dto";

export class TodoDto {
  @ApiProperty()
  @IsString()
  id: string;

  @ApiProperty({ type: UserDto })
  @Type(() => UserDto)
  user: UserDto;

  @ApiProperty()
  @IsString()
  title: string;

  @ApiProperty()
  @IsDateString()
  date: Date;
}