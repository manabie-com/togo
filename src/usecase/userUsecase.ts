import UserRepo from '../repo/user';
import { UserRequest } from '../dto/userRequest';
import { AuthRequest } from '../dto/authRequest';
import { User } from '../entity/User';
import { hash, compare } from '../util/hasher';
import { jwt } from '../util/jwt';
import { Error } from '../error';

const userUsecase = (repo: UserRepo) => ({
	create: async (u: UserRequest) => {
		u.validate();
		const user = new User();
		user.username = u.username;
		user.password = await hash(u.password);
		user.type = u.type;
		const userResponse = await repo.create(user);
		return userResponse;
	},
	authenticate: async (u: AuthRequest) => {
		u.validate();
		const user = new User();
		user.username = u.username;
		user.password = u.password;
		const { username, id, type, password } = await repo.authenticate(user);

		const claims = await compare(u.password, password);

		if (claims) {
			return { token: jwt.generate({ id, username, type }) }
		}

		Error.exec('User not found on database', 401);
		return { token: '' };
	}
})

export default userUsecase;
