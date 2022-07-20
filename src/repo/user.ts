import { User } from '../entity/User';
import { UserResponse } from '../dto/userResponse';

interface UserRepo {
	create(u: User): Promise<UserResponse>
}

export default UserRepo;
