import { UserRequest } from '../dto/userRequest';
import { UserResponse } from '../dto/userResponse';

interface UserUsecase {
	create(u: UserRequest): Promise<UserResponse>
}

export default UserUsecase;

