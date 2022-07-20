import { DataSource } from 'typeorm';
import { User } from '../entity/User';
import userResponse, { UserResponse } from '../dto/userResponse';

const userRepo = (db: DataSource) => ({
	create: async (u: User): Promise<UserResponse> => {
		await db.manager.save(u);
		userResponse.id = u.id;
		userResponse.username = u.username;
		userResponse.type = u.type;
		return userResponse;
	}
})

export default userRepo;
