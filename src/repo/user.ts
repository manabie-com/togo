/*
 * Repository Interface which defines repository signature
 */
import { User } from '../entity/User';
import { UserResponse } from '../dto/userResponse';
import { AuthUserResponse } from '../dto/authResponse';

interface UserRepo {
	create(u: User): Promise<UserResponse>
	authenticate(u: User): Promise<AuthUserResponse>
	upgrade(u: User): Promise<boolean>
}

export default UserRepo;
