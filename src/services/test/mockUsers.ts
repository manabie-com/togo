import { User } from 'src/entities/user.entity';

export const mockedUsers = [
  new User({ id: 1, name: 'Nguyen Van A', dailyMaxTasks: 2 }),
  new User({ id: 2, name: 'Nguyen Van A', dailyMaxTasks: 1 }),
];
