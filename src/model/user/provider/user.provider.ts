import { USER } from "src/constance/variable";
import { User } from "../schema/user.entity";

export const userProviders = [
  {
    provide: USER,
    useValue: User,
  },
];