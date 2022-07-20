import { UserRequest } from '../dto/userRequest';
import { UserResponse } from '../dto/userResponse';
import { AuthResponse } from '../dto/authResponse';
import { AuthRequest } from '../dto/authRequest';

interface UserUsecase {
	create(u: UserRequest): Promise<UserResponse>
	authenticate(u: AuthRequest): Promise<AuthResponse>
}

export default UserUsecase;

