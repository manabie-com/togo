import UserRepo from '../repo/user';
import { UserRequest } from '../dto/userRequest';
import { User } from '../entity/User';

const userUsecase = (repo: UserRepo) => ({
	create: async (u: UserRequest) => {
		const user = new User();
		user.username = u.username;
		user.password = u.password;
		const userResponse = await repo.create(user);
		return userResponse
	}
})

export default userUsecase;
