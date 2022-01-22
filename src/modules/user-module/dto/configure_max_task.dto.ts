import { IsNotEmpty, IsNumber, Min } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class ConfigureMaxTaskDTO {
    @IsNumber()
    @IsNotEmpty()
    @ApiProperty()
    id: number;

    @IsNumber()
    @Min(0)
    @IsNotEmpty()
    @ApiProperty()
    maxTask: number;
}
