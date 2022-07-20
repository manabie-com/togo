import { User } from '../entity/User';
import { UserResponse } from '../dto/userResponse';
import { AuthUserResponse } from '../dto/authResponse';

interface UserRepo {
	create(u: User): Promise<UserResponse>
	authenticate(u: User): Promise<AuthUserResponse>
}

export default UserRepo;
