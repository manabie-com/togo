import { CreateUserDto } from "../../src/User/dto/create-User-dto";
import { User } from "../../src/User/schemas/User.schema";

let createUserDto = new CreateUserDto();
createUserDto.name = 'anyname';
createUserDto.dailyTaskLimit = 5;

export const mockCreateUserDto = (): CreateUserDto => ({
  ...createUserDto
});

export const mockUser = (): User => ({
  id: 'anyid',
  ...createUserDto,
  dailyTaskCounter: 0,
  dailyTaskDate: new Date(),
  createdAt: new Date(),
  updatedAt: new Date()
})