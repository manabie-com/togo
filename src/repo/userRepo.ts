import { DataSource } from 'typeorm';
import { User } from '../entity/User';
import { userResponse } from '../dto/userResponse';
import { authUserResponse } from '../dto/authResponse';
import { Error } from '../error';
import { CONSTANTS } from '../util/constant';

const userRepo = (db: DataSource) => ({
	create: async (u: User) => {
		try {
		await db.manager.save(u);
		userResponse.id = u.id;
		userResponse.username = u.username;
		userResponse.type = u.type;
		return userResponse;
		} catch (e: any) {
			console.log(e);
			Error.exec('Database error', 500);
			return userResponse;
		}
	},
	authenticate: async (u: User) => {
		try {
			const repo = db.getRepository(User);
			const user = await repo.findOneBy({
				username: u.username,
			})

			if (user) {
				authUserResponse.id = user.id;
				authUserResponse.username = user.username;
				authUserResponse.password = user.password;
				authUserResponse.type = user.type;
			}

			return authUserResponse;
		} catch (e: any) {
			console.log(e);
			Error.exec('Database error', 500);
			return authUserResponse;
		}
	},
	upgrade: async (u: User) => {
		try {
			const repo = db.getRepository(User);
			const user = await repo.findOneBy({
				id: u.id,
			});

			if (user) {
				user.type = CONSTANTS.ACCOUNT_TYPE.PREMIUM;
				await repo.save(user);
			}

			return true;

		} catch (e: any) {
			console.log(e);
			Error.exec('Database error', 500);
			return false;
		}
	},
})

export default userRepo;
