import {IsNumber, IsOptional, Min } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';
import { Type } from 'class-transformer';

export class FindUserDTO {
    @IsOptional()
    @IsNumber()
    @ApiProperty({ required: false })
    @Type(() => Number)
    @Min(1)
    pageIndex: number;

    @IsOptional()
    @IsNumber()
    @Type(() => Number)
    @ApiProperty({ required: false })
    @Min(1)
    perPage: number;
}
