/* 
 * DTO - Data Transfer Object
 * Handles data passing from controller to useCase
 * Also Handles input validation
 */
export interface UserResponse {
	id: number;
	username: string;
	type: string;
}

export const userResponse: UserResponse = {
	id: 0,
	username: '',
	type: '',
}
