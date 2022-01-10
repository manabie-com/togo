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


export const userId1 = "305c8624-a214-4a25-93b4-d54dcc411150";
export const userId2 = "d864710e-c343-413d-8885-d2d0053fa75b";
export const userId3 = "7d4c3161-5c7b-4686-86cd-26db29668e44";
export const mockUsersList = (): User[] => ([
  {
    id: userId1,
    name: 'user1',
    dailyTaskLimit: 2,
    dailyTaskCounter: 2,
    dailyTaskDate: new Date(),
    createdAt: new Date(),
    updatedAt: new Date()
  },
  {
    id: userId2,
    name: 'user2',
    dailyTaskLimit: 1,
    dailyTaskCounter: 1,
    dailyTaskDate: new Date("2022-01-09T01:50:12.000Z"),
    createdAt: new Date(),
    updatedAt: new Date()
  },
  {
    id: userId3,
    name: 'user3',
    dailyTaskLimit: 3,
    dailyTaskCounter: 2,
    dailyTaskDate: new Date(),
    createdAt: new Date(),
    updatedAt: new Date()
  }
]);