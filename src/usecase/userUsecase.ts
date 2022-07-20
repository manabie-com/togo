import UserRepo from '../repo/user';
import { UserRequest } from '../dto/userRequest';
import { User } from '../entity/User';
import { hash } from '../util/hasher';

const userUsecase = (repo: UserRepo) => ({
	create: async (u: UserRequest) => {
		u.validate()
		const user = new User();
		user.username = u.username;
		user.password = await hash(u.password);
		user.type = u.type;
		const userResponse = await repo.create(user);
		return userResponse
	}
})

export default userUsecase;
