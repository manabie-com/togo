import { User } from "./schemas/user.schema";

export const userId1 = "305c8624-a214-4a25-93b4-d54dcc411150";
export const userId2 = "d864710e-c343-413d-8885-d2d0053fa75b";
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
    dailyTaskLimit: 3,
    dailyTaskCounter: 1,
    dailyTaskDate: new Date(),
    createdAt: new Date(),
    updatedAt: new Date()
  }
]);