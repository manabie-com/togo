import { DataSource } from 'typeorm';
import { User } from '../entity/User';
import userResponse from '../dto/userResponse';
import { Error } from '../error';

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
	}
})

export default userRepo;
