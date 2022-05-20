import { IsNotEmpty, IsNumber, IsString } from 'class-validator';

export class CreateTaskDto {
    @IsNotEmpty()
    @IsString()
    content: string;

    @IsNotEmpty()
    @IsNumber()
    userId: number;
}